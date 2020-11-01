package ias_zone

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ZoneEnrollResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ZoneEnrollResponse{ResponseCode: 0x01, ZoneID: 0x02}
		actualCommand := ZoneEnrollResponse{}
		expectedBytes := []byte{0x01, 0x02}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ClientToServer, &ZoneEnrollResponse{})
		assert.NoError(t, err)
		assert.Equal(t, ZoneEnrollResponseId, id)
	})
}

func Test_InitiateNormalOperationMode(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := InitiateNormalOperationMode{}
		actualCommand := InitiateNormalOperationMode{}
		var expectedBytes []byte

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

		id, err := cr.GetLocalCommandIdentifier(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ClientToServer, &InitiateNormalOperationMode{})
		assert.NoError(t, err)
		assert.Equal(t, InitiateNormalOperationModeId, id)
	})
}

func Test_InitiateTestMode(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := InitiateTestMode{TestModeDuration: 0x01, CurrentZoneSensitivityLevel: 0x02}
		actualCommand := InitiateTestMode{}
		expectedBytes := []byte{0x01, 0x02}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ClientToServer, &InitiateTestMode{})
		assert.NoError(t, err)
		assert.Equal(t, InitiateTestModeId, id)
	})
}

func Test_ZoneStatusChangeNotification(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ZoneStatusChangeNotification{
			BatteryDefect:      true,
			TestMode:           false,
			ACMainsFault:       false,
			Trouble:            true,
			RestoreReports:     false,
			SupervisionReports: true,
			BatteryLow:         false,
			Tamper:             true,
			Alarm2:             false,
			Alarm1:             true,
			ExtendedStatus:     0xff,
			ZoneID:             0x02,
			Delay:              0x0304,
		}
		actualCommand := ZoneStatusChangeNotification{}
		expectedBytes := []byte{0x02, 0x55, 0xff, 0x02, 0x04, 0x03}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ServerToClient, &ZoneStatusChangeNotification{})
		assert.NoError(t, err)
		assert.Equal(t, ZoneStatusChangeNotificationId, id)
	})
}

func Test_ZoneEnrollRequest(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ZoneEnrollRequest{ZoneType: 0x0102, ManufacturerCode: 0x0304}
		actualCommand := ZoneEnrollRequest{}
		expectedBytes := []byte{0x02, 0x01, 0x04, 0x03}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ServerToClient, &ZoneEnrollRequest{})
		assert.NoError(t, err)
		assert.Equal(t, ZoneEnrollRequestId, id)
	})
}
