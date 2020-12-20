package level

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MoveToLevel(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveToLevel{
			Level:          0xaa,
			TransitionTime: 0x1122,
		}
		actualCommand := MoveToLevel{}
		expectedBytes := []byte{0xaa, 0x22, 0x11}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveToLevel{})
		assert.NoError(t, err)
		assert.Equal(t, MoveToLevelId, id)
	})
}

func Test_Move(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := Move{
			MoveMode: 0x01,
			Rate:     0x02,
		}
		actualCommand := Move{}
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

		id, err := cr.GetLocalCommandIdentifier(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, &Move{})
		assert.NoError(t, err)
		assert.Equal(t, MoveId, id)
	})
}

func Test_Step(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := Step{
			StepMode:       0x01,
			StepSize:       0x02,
			TransitionTime: 0x0304,
		}
		actualCommand := Step{}
		expectedBytes := []byte{0x01, 0x02, 0x04, 0x03}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, &Step{})
		assert.NoError(t, err)
		assert.Equal(t, StepId, id)
	})
}

func Test_Stop(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := Stop{}
		actualCommand := Stop{}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Empty(t, actualBytes)

		err = bytecodec.Unmarshal(actualBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetLocalCommandIdentifier(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, &Stop{})
		assert.NoError(t, err)
		assert.Equal(t, StopId, id)
	})
}

func Test_MoveToLevelWithOnOff(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveToLevelWithOnOff{
			Level:          0xaa,
			TransitionTime: 0x1122,
		}
		actualCommand := MoveToLevelWithOnOff{}
		expectedBytes := []byte{0xaa, 0x22, 0x11}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveToLevelWithOnOff{})
		assert.NoError(t, err)
		assert.Equal(t, MoveToLevelWithOnOffId, id)
	})
}

func Test_MoveWithOnOff(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveWithOnOff{
			MoveMode: 0x01,
			Rate:     0x02,
		}
		actualCommand := MoveWithOnOff{}
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

		id, err := cr.GetLocalCommandIdentifier(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveWithOnOff{})
		assert.NoError(t, err)
		assert.Equal(t, MoveWithOnOffId, id)
	})
}

func Test_StepWithOnOff(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := StepWithOnOff{
			StepMode:       0x01,
			StepSize:       0x02,
			TransitionTime: 0x0304,
		}
		actualCommand := StepWithOnOff{}
		expectedBytes := []byte{0x01, 0x02, 0x04, 0x03}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, &StepWithOnOff{})
		assert.NoError(t, err)
		assert.Equal(t, StepWithOnOffId, id)
	})
}

func Test_StopWithOnOff(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := StopWithOnOff{}
		actualCommand := StopWithOnOff{}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Empty(t, actualBytes)

		err = bytecodec.Unmarshal(actualBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetLocalCommandIdentifier(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, &StopWithOnOff{})
		assert.NoError(t, err)
		assert.Equal(t, StopWithOnOffId, id)
	})
}
