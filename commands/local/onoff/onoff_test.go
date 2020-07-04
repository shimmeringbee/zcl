package onoff

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Off(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := Off{}
		actualCommand := Off{}
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

		id, err := cr.GetLocalCommandIdentifier(zcl.OnOffId, zigbee.NoManufacturer, &Off{})
		assert.NoError(t, err)
		assert.Equal(t, OffId, id)
	})
}

func Test_On(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := On{}
		actualCommand := On{}
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

		id, err := cr.GetLocalCommandIdentifier(zcl.OnOffId, zigbee.NoManufacturer, &On{})
		assert.NoError(t, err)
		assert.Equal(t, OnId, id)
	})
}

func Test_Toggle(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := Toggle{}
		actualCommand := Toggle{}
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

		id, err := cr.GetLocalCommandIdentifier(zcl.OnOffId, zigbee.NoManufacturer, &Toggle{})
		assert.NoError(t, err)
		assert.Equal(t, ToggleId, id)
	})
}

func Test_OffWithEffect(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := OffWithEffect{
			EffectIdentifier: 0x55,
			EffectVariant:    0xaa,
		}
		actualCommand := OffWithEffect{}
		expectedBytes := []byte{0x55, 0xaa}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.OnOffId, zigbee.NoManufacturer, &OffWithEffect{})
		assert.NoError(t, err)
		assert.Equal(t, OffWithEffectId, id)
	})
}

func Test_OnWithRecallGlobalScene(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := OnWithRecallGlobalScene{}
		actualCommand := OnWithRecallGlobalScene{}
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

		id, err := cr.GetLocalCommandIdentifier(zcl.OnOffId, zigbee.NoManufacturer, &OnWithRecallGlobalScene{})
		assert.NoError(t, err)
		assert.Equal(t, OnWithRecallGlobalSceneId, id)
	})
}

func Test_OnWithTimedOff(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := OnWithTimedOff{
			OnOffControl: 0x01,
			OnTime:       0x1122,
			OffWaitTime:  0x3344,
		}
		actualCommand := OnWithTimedOff{}
		expectedBytes := []byte{0x01, 0x22, 0x11, 0x44, 0x33}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.OnOffId, zigbee.NoManufacturer, &OnWithTimedOff{})
		assert.NoError(t, err)
		assert.Equal(t, OnWithTimedOffId, id)
	})
}
