package zcl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Marshal(t *testing.T) {
	t.Run("a manufacturer specific header and message marshals", func(t *testing.T) {
		message := Message{
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

		expectedBytes := []byte{0b00000100, 0x20, 0x10, 0x40, 0x0b, 0xaa, 0x01}
		actualBytes, err := Marshal(message)

		assert.NoError(t, err)
		assert.Equal(t, actualBytes, expectedBytes)
	})

	t.Run("no manufacturer specific header and message marshals", func(t *testing.T) {
		message := Message{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   false,
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

		expectedBytes := []byte{0b0000000, 0x40, 0x0b, 0xaa, 0x01}
		actualBytes, err := Marshal(message)

		assert.NoError(t, err)
		assert.Equal(t, actualBytes, expectedBytes)
	})
}
