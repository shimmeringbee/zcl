package zcl

import (
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Unmarshal(t *testing.T) {
	t.Run("message with unknown command identifier", func(t *testing.T) {
		in := zigbee.ApplicationMessage{
			ClusterID:           0x0102,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000100, 0x20, 0x10, 0x40, 0xff},
		}

		_, err := Unmarshal(in)
		assert.Error(t, err)
	})

	t.Run("DefaultResponse with manufacturer specific", func(t *testing.T) {
		in := zigbee.ApplicationMessage{
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000100, 0x20, 0x10, 0x40, 0x0b, 0xaa, 0x01},
		}

		expectedMessage := ZCLMessage{
			FrameType:           FrameGlobal,
			Direction:           ClientToServer,
			TransactionSequence: 0x40,
			Manufacturer:        0x1020,
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Command: &DefaultResponse{
				CommandIdentifier: 0xaa,
				Status:            0x01,
			},
		}

		actualMessage, err := Unmarshal(in)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DefaultResponse without manufacturer specific", func(t *testing.T) {
		in := zigbee.ApplicationMessage{
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000000, 0x40, 0x0b, 0xaa, 0x01},
		}

		expectedMessage := ZCLMessage{
			FrameType:           FrameGlobal,
			Direction:           ClientToServer,
			TransactionSequence: 0x40,
			Manufacturer:        0x0,
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Command: &DefaultResponse{
				CommandIdentifier: 0xaa,
				Status:            0x01,
			},
		}

		actualMessage, err := Unmarshal(in)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})
}
