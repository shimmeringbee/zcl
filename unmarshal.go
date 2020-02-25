package zcl

import (
	"fmt"
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
)

func Unmarshal(data []byte) (ZCLFrame, error) {
	frame := ZCLFrame{}

	bb := bitbuffer.NewBitBufferFromBytes(data)

	if err := bytecodec.UnmarshalFromBitBuffer(bb, &frame.Header); err != nil {
		return ZCLFrame{}, err
	}

	switch frame.Header.CommandIdentifier {
	case ReadAttributesID:
		frame.Command = &ReadAttributes{}
	case ReadAttributesResponseID:
		frame.Command = &ReadAttributesResponse{}
	case WriteAttributesID:
		frame.Command = &WriteAttributes{}
	case WriteAttributesUndividedID:
		frame.Command = &WriteAttributesUndivided{}
	case WriteAttributesResponseID:
		frame.Command = &WriteAttributesResponse{}
	case WriteAttributesNoResponseID:
		frame.Command = &WriteAttributesNoResponse{}
	case ReportAttributesID:
		frame.Command = &ReportAttributes{}
	case DefaultResponseID:
		frame.Command = &DefaultResponse{}
	case DiscoverAttributesID:
		frame.Command = &DiscoverAttributes{}
	case DiscoverAttributesResponseID:
		frame.Command = &DiscoverAttributesResponse{}
	case ReadAttributesStructuredID:
		frame.Command = &ReadAttributesStructured{}
	case WriteAttributesStructuredID:
		frame.Command = &WriteAttributesStructured{}
	case WriteAttributesStructuredResponseID:
		frame.Command = &WriteAttributesStructuredResponse{}
	case DiscoverCommandsReceivedID:
		frame.Command = &DiscoverCommandsReceived{}
	case DiscoverCommandsReceivedResponseID:
		frame.Command = &DiscoverCommandsReceivedResponse{}
	case DiscoverCommandsGeneratedID:
		frame.Command = &DiscoverCommandsGenerated{}
	case DiscoverCommandsGeneratedResponseID:
		frame.Command = &DiscoverCommandsGeneratedResponse{}
	case DiscoverAttributesExtendedID:
		frame.Command = &DiscoverAttributesExtended{}
	case DiscoverAttributesExtendedResponseID:
		frame.Command = &DiscoverAttributesExtendedResponse{}
	default:
		return ZCLFrame{}, fmt.Errorf("unknown ZCL command identifier received: %d", frame.Header.CommandIdentifier)
	}

	if err := bytecodec.UnmarshalFromBitBuffer(bb, frame.Command); err != nil {
		return ZCLFrame{}, err
	}

	return frame, nil
}
