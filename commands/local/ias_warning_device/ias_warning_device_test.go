package ias_warning_device

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StartWarning(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := StartWarning{
			WarningMode:     Emergency,
			StrobeMode:      StrobeWithWarning,
			WarningDuration: 0x55aa,
		}
		actualCommand := StartWarning{}
		expectedBytes := []byte{0b00110100, 0xaa, 0x55}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetLocalCommandIdentifier(zcl.IASWarningDevicesId, zigbee.NoManufacturer, zcl.ClientToServer, &StartWarning{})
		assert.NoError(t, err)
		assert.Equal(t, StartWarningId, id)
	})
}

func Test_Squawk(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := Squawk{
			SquawkMode:  SystemDisarmed,
			Strobe:      true,
			SquawkLevel: High,
		}
		actualCommand := Squawk{}
		expectedBytes := []byte{0b00011010}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetLocalCommandIdentifier(zcl.IASWarningDevicesId, zigbee.NoManufacturer, zcl.ClientToServer, &Squawk{})
		assert.NoError(t, err)
		assert.Equal(t, SquawkId, id)
	})
}
