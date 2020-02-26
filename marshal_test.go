package zcl

import (
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Marshal(t *testing.T) {
	type Command struct {
		FieldOne uint8
	}

	clusterID := zigbee.ClusterID(0xbeef)
	commandID := CommandIdentifier(0xcc)
	manufacturer := uint16(0x1020)

	cr := NewCommandRegistry()
	cr.RegisterGlobal(commandID, &Command{})
	cr.RegisterLocal(clusterID, manufacturer, commandID, &Command{})

	t.Run("a manufacturer specific header and global message marshals", func(t *testing.T) {
		in := Message{
			FrameType:           FrameGlobal,
			Direction:           ClientToServer,
			TransactionSequence: 0x40,
			Manufacturer:        0x1020,
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Command: &Command{
				FieldOne: 0xaa,
			},
		}

		expectedOut := zigbee.ApplicationMessage{
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000100, 0x20, 0x10, 0x40, 0xcc, 0xaa},
		}

		actualOut, err := Marshal(cr, in)

		assert.NoError(t, err)
		assert.Equal(t, expectedOut, actualOut)
	})

	t.Run("no manufacturer specific header and global message marshals", func(t *testing.T) {
		in := Message{
			FrameType:           FrameGlobal,
			Direction:           ClientToServer,
			TransactionSequence: 0x40,
			Manufacturer:        0x0,
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Command: &Command{
				FieldOne: 0xaa,
			},
		}

		expectedOut := zigbee.ApplicationMessage{
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000000, 0x40, 0xcc, 0xaa},
		}

		actualOut, err := Marshal(cr, in)

		assert.NoError(t, err)
		assert.Equal(t, expectedOut, actualOut)
	})

	t.Run("a manufacturer specific header and a local message marshals", func(t *testing.T) {
		in := Message{
			FrameType:           FrameLocal,
			Direction:           ClientToServer,
			TransactionSequence: 0x40,
			Manufacturer:        manufacturer,
			ClusterID:           clusterID,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Command: &Command{
				FieldOne: 0xaa,
			},
		}

		expectedOut := zigbee.ApplicationMessage{
			ClusterID:           clusterID,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000101, 0x20, 0x10, 0x40, 0xcc, 0xaa},
		}

		actualOut, err := Marshal(cr, in)

		assert.NoError(t, err)
		assert.Equal(t, expectedOut, actualOut)
	})
}
