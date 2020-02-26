package zcl

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Header(t *testing.T) {
	t.Run("manufacturer specific bit includes manufacturer ID in marshalled frame", func(t *testing.T) {
		frame := Header{
			Control: Control{
				DisableDefaultResponse: false,
				Direction:              ClientToServer,
				ManufacturerSpecific:   true,
				FrameType:              FrameLocal,
			},
			Manufacturer:        0xaabb,
			TransactionSequence: 0xcc,
			CommandIdentifier:   0xdd,
		}

		data, err := bytecodec.Marshal(&frame)

		assert.NoError(t, err)
		assert.Equal(t, []byte{0b00000101, 0xbb, 0xaa, 0xcc, 0xdd}, data)
	})

	t.Run("absent manufacturer specific bit excludes manufacturer ID in marshalled frame", func(t *testing.T) {
		frame := Header{
			Control: Control{
				DisableDefaultResponse: false,
				Direction:              ClientToServer,
				ManufacturerSpecific:   false,
				FrameType:              FrameLocal,
			},
			Manufacturer:        0xaabb,
			TransactionSequence: 0xcc,
			CommandIdentifier:   0xdd,
		}

		data, err := bytecodec.Marshal(&frame)

		assert.NoError(t, err)
		assert.Equal(t, []byte{0b00000001, 0xcc, 0xdd}, data)
	})
}

func Test_ZCLMessage(t *testing.T) {
	t.Run("a message without a manufacturer identifier is not identified as manufacturer specific", func(t *testing.T) {
		message := Message{}
		assert.False(t, message.isManufacturerSpecific())
	})
	t.Run("a message with a manufacturer identifier is identified as manufacturer specific", func(t *testing.T) {
		message := Message{Manufacturer: 0x0001}
		assert.True(t, message.isManufacturerSpecific())
	})
}
