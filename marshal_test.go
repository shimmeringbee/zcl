package zcl

import (
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Marshal(t *testing.T) {
	t.Run("a manufacturer specific header and message marshals", func(t *testing.T) {
		in := ZCLMessage{
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

		expectedOut := zigbee.ApplicationMessage{
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000100, 0x20, 0x10, 0x40, 0x0b, 0xaa, 0x01},
		}

		actualOut, err := Marshal(in)

		assert.NoError(t, err)
		assert.Equal(t, expectedOut, actualOut)
	})

	t.Run("no manufacturer specific header and message marshals", func(t *testing.T) {
		in := ZCLMessage{
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

		expectedOut := zigbee.ApplicationMessage{
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000000, 0x40, 0x0b, 0xaa, 0x01},
		}

		actualOut, err := Marshal(in)

		assert.NoError(t, err)
		assert.Equal(t, expectedOut, actualOut)
	})
}
