package zcl

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
	"github.com/shimmeringbee/zigbee"
)

func Marshal(message ZCLMessage) (zigbee.ApplicationMessage, error) {
	bb := bitbuffer.NewBitBuffer()

	header := Header{
		Control: Control{
			Reserved:               0,
			DisableDefaultResponse: false,
			Direction:              message.Direction,
			ManufacturerSpecific:   message.isManufacturerSpecific(),
			FrameType:              message.FrameType,
		},
		Manufacturer:        message.Manufacturer,
		TransactionSequence: message.TransactionSequence,
		CommandIdentifier:   0x0,
	}

	if err := bytecodec.MarshalToBitBuffer(bb, header); err != nil {
		return zigbee.ApplicationMessage{}, err
	}

	if err := bytecodec.MarshalToBitBuffer(bb, message.Command); err != nil {
		return zigbee.ApplicationMessage{}, err
	}

	msg := zigbee.ApplicationMessage{
		ClusterID:           message.ClusterID,
		SourceEndpoint:      message.SourceEndpoint,
		DestinationEndpoint: message.DestinationEndpoint,
		Data:                bb.Bytes(),
	}

	return msg, nil
}
