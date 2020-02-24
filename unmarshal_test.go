package zcl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Unmarshal(t *testing.T) {
	t.Run("message with unknown command identifier", func(t *testing.T) {
		data := []byte{0b00000100, 0x20, 0x10, 0x40, 0xff}

		_, err := Unmarshal(data)
		assert.Error(t, err)
	})

	t.Run("DefaultResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := Message{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DefaultResponseID,
			},
			Command: &DefaultResponse{
				CommandIdentifier: 0xaa,
				Status:            0x01,
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})
}
