package zcl

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
)

func Marshal(message Message) ([]byte, error) {
	bb := bitbuffer.NewBitBuffer()

	if err := bytecodec.MarshalToBitBuffer(bb, message.Header); err != nil {
		return []byte{}, err
	}

	if err := bytecodec.MarshalToBitBuffer(bb, message.Command); err != nil {
		return []byte{}, err
	}

	return bb.Bytes(), nil
}
