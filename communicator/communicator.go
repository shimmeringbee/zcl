package communicator

import (
	"context"
	"errors"
	"fmt"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zcl/commands/global"
	"github.com/shimmeringbee/zigbee"
	"sync"
	"sync/atomic"
)

type MessageWithSource struct {
	SourceAddress zigbee.IEEEAddress
	Message       zcl.Message
}

var matchId = new(uint64)

type Matcher func(address zigbee.IEEEAddress, appMsg zigbee.ApplicationMessage, zclMessage zcl.Message) bool

func AddressAndSequenceMatch(matchAddress zigbee.IEEEAddress, matchSequence uint8) Matcher {
	return func(address zigbee.IEEEAddress, appMsg zigbee.ApplicationMessage, zclMessage zcl.Message) bool {
		return matchAddress == address && matchSequence == zclMessage.TransactionSequence
	}
}

func NewMatch(matcher Matcher, callback func(source MessageWithSource)) Match {
	return Match{
		id:       atomic.AddUint64(matchId, 1),
		matcher:  matcher,
		callback: callback,
	}
}

type Match struct {
	id       uint64
	matcher  Matcher
	callback func(source MessageWithSource)
}

type communicator struct {
	Provider        zigbee.Provider
	CommandRegistry *zcl.CommandRegistry

	mutex   *sync.RWMutex
	matches map[uint64]Match
}

func NewCommunicator(provider zigbee.Provider, registry *zcl.CommandRegistry) Communicator {
	return &communicator{
		Provider:        provider,
		CommandRegistry: registry,
		mutex:           &sync.RWMutex{},
		matches:         map[uint64]Match{},
	}
}

func (c *communicator) ProcessIncomingMessage(msg zigbee.NodeIncomingMessageEvent) error {
	message, err := c.CommandRegistry.Unmarshal(msg.ApplicationMessage)

	if err != nil {
		return fmt.Errorf("failed to unmarshal incomming ZCL message: %w", err)
	}

	c.mutex.RLock()
	defer c.mutex.RUnlock()

	for _, match := range c.matches {
		if match.matcher(msg.IEEEAddress, msg.ApplicationMessage, message) {
			go match.callback(MessageWithSource{
				SourceAddress: msg.IEEEAddress,
				Message:       message,
			})
		}
	}

	return nil
}

func (c *communicator) RegisterMatch(match Match) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.matches[match.id] = match
}

func (c *communicator) UnregisterMatch(match Match) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.matches, match.id)
}

func (c *communicator) Request(ctx context.Context, address zigbee.IEEEAddress, requireAck bool, message zcl.Message) error {
	appMessage, err := c.CommandRegistry.Marshal(message)

	if err != nil {
		return fmt.Errorf("ZCL communicator failed to send message during marshalling: %w", err)
	}

	err = c.Provider.SendApplicationMessageToNode(ctx, address, appMessage, requireAck)

	if err != nil {
		return fmt.Errorf("ZCL communicator failed to send via provider: %w", err)
	}

	return nil
}

func (c *communicator) RequestResponse(ctx context.Context, address zigbee.IEEEAddress, requireAck bool, message zcl.Message) (zcl.Message, error) {
	ch := make(chan zcl.Message, 1)

	match := NewMatch(AddressAndSequenceMatch(address, message.TransactionSequence),
		func(recvMessage MessageWithSource) {
			ch <- recvMessage.Message
		})

	c.RegisterMatch(match)
	defer c.UnregisterMatch(match)

	if err := c.Request(ctx, address, requireAck, message); err != nil {
		return zcl.Message{}, err
	}

	select {
	case resp := <-ch:
		return resp, nil
	case <-ctx.Done():
		return zcl.Message{}, errors.New("ZCL communicator waiting for reply, context expired")
	}
}

func (c *communicator) ReadAttributes(ctx context.Context, ieeeAddress zigbee.IEEEAddress, requireAck bool, cluster zigbee.ClusterID, code zigbee.ManufacturerCode, sourceEndpoint zigbee.Endpoint, destEndpoint zigbee.Endpoint, transactionSequence uint8, attributes []zcl.AttributeID) ([]global.ReadAttributeResponseRecord, error) {
	request := zcl.Message{
		FrameType:           zcl.FrameGlobal,
		Direction:           zcl.ClientToServer,
		TransactionSequence: transactionSequence,
		Manufacturer:        code,
		ClusterID:           cluster,
		SourceEndpoint:      sourceEndpoint,
		DestinationEndpoint: destEndpoint,
		Command: &global.ReadAttributes{
			Identifier: attributes,
		},
	}

	response, err := c.RequestResponse(ctx, ieeeAddress, requireAck, request)

	if err != nil {
		return nil, err
	}

	if readResponse, is := response.Command.(*global.ReadAttributesResponse); is {
		return readResponse.Records, nil
	} else {
		return []global.ReadAttributeResponseRecord{}, errors.New("read attributes received command back which was not ReadAttributesResponse")
	}
}

func ReadResponsesToMap(recs []global.ReadAttributeResponseRecord) map[zcl.AttributeID]global.ReadAttributeResponseRecord {
	m := map[zcl.AttributeID]global.ReadAttributeResponseRecord{}

	for _, r := range recs {
		m[r.Identifier] = r
	}

	return m
}

func (c *communicator) WriteAttributes(ctx context.Context, ieeeAddress zigbee.IEEEAddress, requireAck bool, cluster zigbee.ClusterID, code zigbee.ManufacturerCode, sourceEndpoint zigbee.Endpoint, destEndpoint zigbee.Endpoint, transactionSequence uint8, attributes map[zcl.AttributeID]zcl.AttributeDataTypeValue) ([]global.WriteAttributesResponseRecord, error) {
	var records []global.WriteAttributesRecord

	for k, v := range attributes {
		records = append(records, global.WriteAttributesRecord{
			Identifier:    k,
			DataTypeValue: &v,
		})
	}

	request := zcl.Message{
		FrameType:           zcl.FrameGlobal,
		Direction:           zcl.ClientToServer,
		TransactionSequence: transactionSequence,
		Manufacturer:        code,
		ClusterID:           cluster,
		SourceEndpoint:      sourceEndpoint,
		DestinationEndpoint: destEndpoint,
		Command: &global.WriteAttributes{
			Records: records,
		},
	}

	response, err := c.RequestResponse(ctx, ieeeAddress, requireAck, request)

	if err != nil {
		return nil, err
	}
	if resp, is := response.Command.(*global.WriteAttributesResponse); is {
		return resp.Records, nil
	} else {
		return nil, errors.New("read attributes received command back which was not WriteAttributesResponse")
	}
}

func WriteResponsesToMap(recs []global.WriteAttributesResponseRecord) map[zcl.AttributeID]global.WriteAttributesResponseRecord {
	m := map[zcl.AttributeID]global.WriteAttributesResponseRecord{}

	for _, r := range recs {
		m[r.Identifier] = r
	}

	return m
}

func (c *communicator) ConfigureReporting(ctx context.Context, ieeeAddress zigbee.IEEEAddress, requireAck bool, cluster zigbee.ClusterID, code zigbee.ManufacturerCode, sourceEndpoint zigbee.Endpoint, destEndpoint zigbee.Endpoint, transactionSequence uint8, attributeId zcl.AttributeID, dataType zcl.AttributeDataType, minimumReportingInterval uint16, maximumReportingInterval uint16, reportableChange interface{}) error {
	request := zcl.Message{
		FrameType:           zcl.FrameGlobal,
		Direction:           zcl.ClientToServer,
		TransactionSequence: transactionSequence,
		Manufacturer:        code,
		ClusterID:           cluster,
		SourceEndpoint:      sourceEndpoint,
		DestinationEndpoint: destEndpoint,
		Command: &global.ConfigureReporting{
			Records: []global.ConfigureReportingRecord{
				{
					Direction:        0x00,
					Identifier:       attributeId,
					DataType:         dataType,
					MinimumInterval:  minimumReportingInterval,
					MaximumInterval:  maximumReportingInterval,
					ReportableChange: &zcl.AttributeDataValue{Value: reportableChange},
					Timeout:          0,
				},
			},
		},
	}

	response, err := c.RequestResponse(ctx, ieeeAddress, requireAck, request)

	if err != nil {
		return err
	}

	if readResponse, is := response.Command.(*global.ConfigureReportingResponse); is {
		if len(readResponse.Records) == 0 {
			return nil
		}

		if readResponse.Records[0].Identifier != attributeId {
			return errors.New("incorrect attribute id response sent to configure reporting")
		}

		if readResponse.Records[0].Status != 0 {
			return errors.New("non success response sent to configure reporting")
		}

		return nil
	} else {
		return errors.New("configure reporting received command back which was not ConfigureReportingResponse")
	}
}
