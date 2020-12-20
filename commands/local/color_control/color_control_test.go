package color_control

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MoveToHue(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveToHue{
			Hue:            0xaa,
			Direction:      DirectionLongestDistance,
			TransitionTime: 0x0203,
		}
		actualCommand := MoveToHue{}
		expectedBytes := []byte{0xaa, 0x01, 0x03, 0x02}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveToHue{})
		assert.NoError(t, err)
		assert.Equal(t, MoveToHueId, id)
	})
}

func Test_MoveHue(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveHue{
			MoveMode: MoveModeDown,
			Rate:     0x55,
		}
		actualCommand := MoveHue{}
		expectedBytes := []byte{0x03, 0x55}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveHue{})
		assert.NoError(t, err)
		assert.Equal(t, MoveHueId, id)
	})
}

func Test_StepHue(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := StepHue{
			StepMode:       StepModeDown,
			StepSize:       0x05,
			TransitionTime: 0x06,
		}
		actualCommand := StepHue{}
		expectedBytes := []byte{0x03, 0x05, 0x06}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &StepHue{})
		assert.NoError(t, err)
		assert.Equal(t, StepHueId, id)
	})
}

func Test_MoveToSaturation(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveToSaturation{
			Saturation:     0x55,
			TransitionTime: 0x0203,
		}
		actualCommand := MoveToSaturation{}
		expectedBytes := []byte{0x55, 0x03, 0x02}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveToSaturation{})
		assert.NoError(t, err)
		assert.Equal(t, MoveToSaturationId, id)
	})
}

func Test_MoveSaturation(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveSaturation{
			MoveMode: MoveModeDown,
			Rate:     0xaa,
		}
		actualCommand := MoveSaturation{}
		expectedBytes := []byte{0x03, 0xaa}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveSaturation{})
		assert.NoError(t, err)
		assert.Equal(t, MoveSaturationId, id)
	})
}

func Test_StepSaturation(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := StepSaturation{
			StepMode:       StepModeDown,
			StepSize:       0x02,
			TransitionTime: 0x04,
		}
		actualCommand := StepSaturation{}
		expectedBytes := []byte{0x03, 0x02, 0x04}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &StepSaturation{})
		assert.NoError(t, err)
		assert.Equal(t, StepSaturationId, id)
	})
}

func Test_MoveToHueAndSaturation(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveToHueAndSaturation{
			Hue:            0x01,
			Saturation:     0x02,
			TransitionTime: 0x0304,
		}
		actualCommand := MoveToHueAndSaturation{}
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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveToHueAndSaturation{})
		assert.NoError(t, err)
		assert.Equal(t, MoveToHueAndSaturationId, id)
	})
}

func Test_MoveToColor(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveToColor{
			ColorX:         0x0102,
			ColorY:         0x0304,
			TransitionTime: 0x0506,
		}
		actualCommand := MoveToColor{}
		expectedBytes := []byte{0x02, 0x01, 0x04, 0x03, 0x06, 0x05}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveToColor{})
		assert.NoError(t, err)
		assert.Equal(t, MoveToColorId, id)
	})
}

func Test_MoveColor(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		t.Skip("bytecodec does not currently support signed ints, see issue #13.")
		expectedCommand := MoveColor{
			RateX: 0x0102,
			RateY: 0x0304,
		}
		actualCommand := MoveColor{}
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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveColor{})
		assert.NoError(t, err)
		assert.Equal(t, MoveColorId, id)
	})
}

func Test_StepColor(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		t.Skip("bytecodec does not currently support signed ints, see issue #13.")
		expectedCommand := StepColor{
			StepX:          0x0102,
			StepY:          0x0304,
			TransitionTime: 0x0506,
		}
		actualCommand := StepColor{}
		expectedBytes := []byte{0x02, 0x01, 0x04, 0x03, 0x06, 0x05}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &StepColor{})
		assert.NoError(t, err)
		assert.Equal(t, StepColorId, id)
	})
}

func Test_MoveToColorTemperature(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveToColorTemperature{
			ColorTemperatureMireds: 0x0102,
			TransitionTime:         0x0304,
		}
		actualCommand := MoveToColorTemperature{}
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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveToColorTemperature{})
		assert.NoError(t, err)
		assert.Equal(t, MoveToColorTemperatureId, id)
	})
}

func Test_EnhancedMoveToHue(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := EnhancedMoveToHue{
			EnhancedHue:    0x0102,
			Direction:      DirectionLongestDistance,
			TransitionTime: 0x0304,
		}
		actualCommand := EnhancedMoveToHue{}
		expectedBytes := []byte{0x02, 0x01, 0x01, 0x04, 0x03}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &EnhancedMoveToHue{})
		assert.NoError(t, err)
		assert.Equal(t, EnhancedMoveToHueId, id)
	})
}

func Test_EnhancedMoveHue(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := EnhancedMoveHue{
			MoveMode: MoveModeDown,
			Rate:     0x0102,
		}
		actualCommand := EnhancedMoveHue{}
		expectedBytes := []byte{0x03, 0x02, 0x01}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &EnhancedMoveHue{})
		assert.NoError(t, err)
		assert.Equal(t, EnhancedMoveHueId, id)
	})
}

func Test_EnhancedStepHue(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := EnhancedStepHue{
			StepMode:       StepModeDown,
			StepSize:       0x0102,
			TransitionTime: 0x0405,
		}
		actualCommand := EnhancedStepHue{}
		expectedBytes := []byte{0x03, 0x02, 0x01, 0x05, 0x04}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &EnhancedStepHue{})
		assert.NoError(t, err)
		assert.Equal(t, EnhancedStepHueId, id)
	})
}

func Test_EnhancedMoveToHueAndSaturation(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := EnhancedMoveToHueAndSaturation{
			EnhancedHue:    0x0102,
			Saturation:     0x03,
			TransitionTime: 0x0405,
		}
		actualCommand := EnhancedMoveToHueAndSaturation{}
		expectedBytes := []byte{0x02, 0x01, 0x03, 0x05, 0x04}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &EnhancedMoveToHueAndSaturation{})
		assert.NoError(t, err)
		assert.Equal(t, EnhancedMoveToHueAndSaturationId, id)
	})
}

func Test_ColorLoopSet(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := ColorLoopSet{
			UpdateStartHue:  true,
			UpdateTime:      true,
			UpdateDirection: true,
			UpdateAction:    true,
			Action:          ActivateFromEnhancedCurrentHue,
			Direction:       Increment,
			Time:            0x0a0b,
			StartHue:        0x0c0d,
		}
		actualCommand := ColorLoopSet{}
		expectedBytes := []byte{0b00001111, 0x02, 0x01, 0x0b, 0x0a, 0x0d, 0x0c}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &ColorLoopSet{})
		assert.NoError(t, err)
		assert.Equal(t, ColorLoopSetId, id)
	})
}

func Test_MoveColorTemperature(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := MoveColorTemperature{
			MoveMode:                      MoveModeDown,
			Rate:                          0x0102,
			ColorTemperatureMinimumMireds: 0x0a0b,
			ColorTemperatureMaximumMireds: 0x0c0d,
		}
		actualCommand := MoveColorTemperature{}
		expectedBytes := []byte{0x03, 0x02, 0x01, 0x0b, 0x0a, 0x0d, 0x0c}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &MoveColorTemperature{})
		assert.NoError(t, err)
		assert.Equal(t, MoveColorTemperatureId, id)
	})
}

func Test_StepColorTemperature(t *testing.T) {
	t.Run("marshals and unmarshalls correctly", func(t *testing.T) {
		expectedCommand := StepColorTemperature{
			StepMode:                      StepModeDown,
			StepSize:                      0x0506,
			TransitionTime:                0xdd55,
			ColorTemperatureMinimumMireds: 0x0a0b,
			ColorTemperatureMaximumMireds: 0x0c0d,
		}
		actualCommand := StepColorTemperature{}
		expectedBytes := []byte{0x03, 0x06, 0x05, 0x55, 0xdd, 0x0b, 0x0a, 0x0d, 0x0c}

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

		id, err := cr.GetLocalCommandIdentifier(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, &StepColorTemperature{})
		assert.NoError(t, err)
		assert.Equal(t, StepColorTemperatureId, id)
	})
}
