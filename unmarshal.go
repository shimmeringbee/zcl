package zcl

import (
	"fmt"
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
)

func Unmarshal(data []byte) (Message, error) {
	message := Message{}

	bb := bitbuffer.NewBitBufferFromBytes(data)

	if err := bytecodec.UnmarshalFromBitBuffer(bb, &message.Header); err != nil {
		return Message{}, err
	}

	switch message.Header.CommandIdentifier {
	case DefaultResponseID:
		message.Command = &DefaultResponse{}
	default:
		return Message{}, fmt.Errorf("unknown ZCL command identifier received: %d", message.Header.CommandIdentifier)
	}

	if err := bytecodec.UnmarshalFromBitBuffer(bb, message.Command); err != nil {
		return Message{}, err
	}

	return message, nil
}
