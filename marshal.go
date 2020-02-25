package zcl

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
)

func Marshal(frame ZCLFrame) ([]byte, error) {
	bb := bitbuffer.NewBitBuffer()

	if err := bytecodec.MarshalToBitBuffer(bb, frame.Header); err != nil {
		return []byte{}, err
	}

	if err := bytecodec.MarshalToBitBuffer(bb, frame.Command); err != nil {
		return []byte{}, err
	}

	return bb.Bytes(), nil
}
