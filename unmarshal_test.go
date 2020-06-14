package zcl

import (
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Unmarshal(t *testing.T) {
	type Command struct {
		FieldOne uint8
	}

	clusterID := zigbee.ClusterID(0xbeef)
	commandID := CommandIdentifier(0xcc)
	manufacturer := uint16(0x1020)

	cr := NewCommandRegistry()
	cr.RegisterGlobal(commandID, &Command{})
	cr.RegisterLocal(clusterID, manufacturer, commandID, &Command{})

	t.Run("message with unknown command identifier", func(t *testing.T) {
		in := zigbee.ApplicationMessage{
			ClusterID:           0x0102,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000100, 0x20, 0x10, 0x40, 0xff},
		}

		_, err := cr.Unmarshal(in)
		assert.Error(t, err)
	})

	t.Run("global Command with manufacturer specific", func(t *testing.T) {
		in := zigbee.ApplicationMessage{
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000100, 0x20, 0x10, 0x40, 0xcc, 0xaa},
		}

		expectedMessage := Message{
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

		actualMessage, err := cr.Unmarshal(in)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("global Command without manufacturer specific", func(t *testing.T) {
		in := zigbee.ApplicationMessage{
			ClusterID:           0x8888,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000000, 0x40, 0xcc, 0xaa},
		}

		expectedMessage := Message{
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

		actualMessage, err := cr.Unmarshal(in)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("Local command with manufacturer specific", func(t *testing.T) {
		in := zigbee.ApplicationMessage{
			ClusterID:           clusterID,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Data:                []byte{0b00000101, 0x20, 0x10, 0x40, byte(commandID), 0xaa},
		}

		expectedMessage := Message{
			FrameType:           FrameLocal,
			Direction:           ClientToServer,
			TransactionSequence: 0x40,
			Manufacturer:        0x1020,
			ClusterID:           clusterID,
			SourceEndpoint:      0x03,
			DestinationEndpoint: 0x04,
			Command: &Command{
				FieldOne: 0xaa,
			},
		}

		actualMessage, err := cr.Unmarshal(in)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})
}
