package zcl

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
	"github.com/shimmeringbee/zigbee"
)

func Marshal(frame ZCLFrame) (zigbee.ApplicationMessage, error) {
	bb := bitbuffer.NewBitBuffer()

	if err := bytecodec.MarshalToBitBuffer(bb, frame.Header); err != nil {
		return zigbee.ApplicationMessage{}, err
	}

	if err := bytecodec.MarshalToBitBuffer(bb, frame.Command); err != nil {
		return zigbee.ApplicationMessage{}, err
	}

	return zigbee.ApplicationMessage{Data: bb.Bytes()}, nil
}
