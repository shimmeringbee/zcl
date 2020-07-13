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

type Matcher func(address zigbee.IEEEAddress, appMsg zigbee.ApplicationMessage, zclMessage zcl.Message) bool

func AddressAndSequenceMatch(matchAddress zigbee.IEEEAddress, matchSequence uint8) Matcher {
	return func(address zigbee.IEEEAddress, appMsg zigbee.ApplicationMessage, zclMessage zcl.Message) bool {
		return matchAddress == address && matchSequence == zclMessage.TransactionSequence
	}
}

func (c *Communicator) NewMatch(matcher Matcher, callback func(source MessageWithSource)) Match {
	return Match{
		Id:       atomic.AddUint64(c.matchId, 1),
		Matcher:  matcher,
		Callback: callback,
	}
}

type Match struct {
	Id       uint64
	Matcher  Matcher
	Callback func(source MessageWithSource)
}

type Communicator struct {
	Provider        zigbee.Provider
	CommandRegistry *zcl.CommandRegistry

	mutex   *sync.RWMutex
	matches map[uint64]Match
	matchId *uint64
}

func NewCommunicator(provider zigbee.Provider, registry *zcl.CommandRegistry) *Communicator {
	return &Communicator{
		Provider:        provider,
		CommandRegistry: registry,
		mutex:           &sync.RWMutex{},
		matches:         map[uint64]Match{},
		matchId:         new(uint64),
	}
}

func (c *Communicator) ProcessIncomingMessage(msg zigbee.NodeIncomingMessageEvent) error {
	message, err := c.CommandRegistry.Unmarshal(msg.ApplicationMessage)

	if err != nil {
		return fmt.Errorf("failed to unmarshal incomming ZCL message: %w", err)
	}

	c.mutex.RLock()
	ourMatches := c.matches
	c.mutex.RUnlock()

	for _, match := range ourMatches {
		if match.Matcher(msg.IEEEAddress, msg.ApplicationMessage, message) {
			go match.Callback(MessageWithSource{
				SourceAddress: msg.IEEEAddress,
				Message:       message,
			})
		}
	}

	return nil
}

func (c *Communicator) AddCallback(match Match) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.matches[match.Id] = match
}

func (c *Communicator) RemoveCallback(match Match) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.matches, match.Id)
}

func (c *Communicator) Request(ctx context.Context, address zigbee.IEEEAddress, requireAck bool, message zcl.Message) error {
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

func (c *Communicator) RequestResponse(ctx context.Context, address zigbee.IEEEAddress, requireAck bool, message zcl.Message) (zcl.Message, error) {
	ch := make(chan zcl.Message, 1)

	match := c.NewMatch(AddressAndSequenceMatch(address, message.TransactionSequence),
		func(recvMessage MessageWithSource) {
			ch <- recvMessage.Message
		})

	c.AddCallback(match)
	defer c.RemoveCallback(match)

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

func (c *Communicator) Global() *GlobalCommunicator {
	return &GlobalCommunicator{communicator: c}
}

type GlobalCommunicator struct {
	communicator *Communicator
}

func (g *GlobalCommunicator) ReadAttributes(ctx context.Context, ieeeAddress zigbee.IEEEAddress, requireAck bool, cluster zigbee.ClusterID, code zigbee.ManufacturerCode, sourceEndpoint zigbee.Endpoint, destEndpoint zigbee.Endpoint, transactionSequence uint8, attributes []zcl.AttributeID) ([]global.ReadAttributeResponseRecord, error) {
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

	response, err := g.communicator.RequestResponse(ctx, ieeeAddress, requireAck, request)

	if err != nil {
		return nil, err
	}

	if readResponse, is := response.Command.(*global.ReadAttributesResponse); is {
		return readResponse.Records, nil
	} else {
		return []global.ReadAttributeResponseRecord{}, errors.New("read attributes received command back which was not ReadAttributesResponse")
	}
}

func (g *GlobalCommunicator) ConfigureReporting(ctx context.Context, ieeeAddress zigbee.IEEEAddress, requireAck bool, cluster zigbee.ClusterID, code zigbee.ManufacturerCode, sourceEndpoint zigbee.Endpoint, destEndpoint zigbee.Endpoint, transactionSequence uint8, attributeId zcl.AttributeID, dataType zcl.AttributeDataType, minimumReportingInterval uint16, maximumReportingInterval uint16, reportableChange interface{}) error {
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

	response, err := g.communicator.RequestResponse(ctx, ieeeAddress, requireAck, request)

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
