package zcl

import (
	"errors"
	"fmt"
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
	"github.com/shimmeringbee/zigbee"
)

func Unmarshal(appMsg zigbee.ApplicationMessage) (ZCLMessage, error) {
	header := Header{}
	var command interface{}

	bb := bitbuffer.NewBitBufferFromBytes(appMsg.Data)

	if err := bytecodec.UnmarshalFromBitBuffer(bb, &header); err != nil {
		return ZCLMessage{}, err
	}

	if header.Control.FrameType != FrameGlobal {
		return ZCLMessage{}, errors.New("can not currently handle any frame which is not a global command")
	}

	switch header.CommandIdentifier {
	case ReadAttributesID:
		command = &ReadAttributes{}
	case ReadAttributesResponseID:
		command = &ReadAttributesResponse{}
	case WriteAttributesID:
		command = &WriteAttributes{}
	case WriteAttributesUndividedID:
		command = &WriteAttributesUndivided{}
	case WriteAttributesResponseID:
		command = &WriteAttributesResponse{}
	case WriteAttributesNoResponseID:
		command = &WriteAttributesNoResponse{}
	case ReportAttributesID:
		command = &ReportAttributes{}
	case DefaultResponseID:
		command = &DefaultResponse{}
	case DiscoverAttributesID:
		command = &DiscoverAttributes{}
	case DiscoverAttributesResponseID:
		command = &DiscoverAttributesResponse{}
	case ReadAttributesStructuredID:
		command = &ReadAttributesStructured{}
	case WriteAttributesStructuredID:
		command = &WriteAttributesStructured{}
	case WriteAttributesStructuredResponseID:
		command = &WriteAttributesStructuredResponse{}
	case DiscoverCommandsReceivedID:
		command = &DiscoverCommandsReceived{}
	case DiscoverCommandsReceivedResponseID:
		command = &DiscoverCommandsReceivedResponse{}
	case DiscoverCommandsGeneratedID:
		command = &DiscoverCommandsGenerated{}
	case DiscoverCommandsGeneratedResponseID:
		command = &DiscoverCommandsGeneratedResponse{}
	case DiscoverAttributesExtendedID:
		command = &DiscoverAttributesExtended{}
	case DiscoverAttributesExtendedResponseID:
		command = &DiscoverAttributesExtendedResponse{}
	default:
		return ZCLMessage{}, fmt.Errorf("unknown ZCL command identifier received: %d", header.CommandIdentifier)
	}

	if err := bytecodec.UnmarshalFromBitBuffer(bb, command); err != nil {
		return ZCLMessage{}, err
	}

	return ZCLMessage{
		FrameType:           header.Control.FrameType,
		Direction:           header.Control.Direction,
		TransactionSequence: header.TransactionSequence,
		Manufacturer:        header.Manufacturer,
		ClusterID:           appMsg.ClusterID,
		SourceEndpoint:      appMsg.SourceEndpoint,
		DestinationEndpoint: appMsg.DestinationEndpoint,
		Command:             command,
	}, nil
}
