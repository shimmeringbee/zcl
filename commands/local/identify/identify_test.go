package identify

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Identify(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := Identify{
			IdentifyTime: 0x1122,
		}
		actualCommand := Identify{}
		expectedBytes := []byte{0x22, 0x11}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.IdentifyId, zigbee.NoManufacturer, zcl.ClientToServer, &Identify{})
		assert.NoError(t, err)
		assert.Equal(t, IdentifyId, id)
	})
}

func Test_IdentifyQuery(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := IdentifyQuery{}
		actualCommand := IdentifyQuery{}
		expectedBytes := []byte(nil)

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

		id, err := cr.GetLocalCommandIdentifier(zcl.IdentifyId, zigbee.NoManufacturer, zcl.ClientToServer, &IdentifyQuery{})
		assert.NoError(t, err)
		assert.Equal(t, IdentifyQueryId, id)
	})
}

func Test_TriggerEffect(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := TriggerEffect{
			EffectIdentifier: 0x11,
			EffectVariant:    0x22,
		}
		actualCommand := TriggerEffect{}
		expectedBytes := []byte{0x11, 0x22}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.IdentifyId, zigbee.NoManufacturer, zcl.ClientToServer, &TriggerEffect{})
		assert.NoError(t, err)
		assert.Equal(t, TriggerEffectId, id)
	})
}

func Test_IdentifyQueryResponse(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := IdentifyQueryResponse{
			Timeout: 0x1122,
		}
		actualCommand := IdentifyQueryResponse{}
		expectedBytes := []byte{0x22, 0x11}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.IdentifyId, zigbee.NoManufacturer, zcl.ServerToClient, &IdentifyQueryResponse{})
		assert.NoError(t, err)
		assert.Equal(t, IdentifyQueryResponseId, id)
	})
}
