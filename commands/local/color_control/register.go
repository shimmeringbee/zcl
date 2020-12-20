package color_control

import (
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
)

func Register(cr *zcl.CommandRegistry) {
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveToHueId, &MoveToHue{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveHueId, &MoveHue{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, StepHueId, &StepHue{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveToSaturationId, &MoveToSaturation{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveSaturationId, &MoveSaturation{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, StepSaturationId, &StepSaturation{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveToHueAndSaturationId, &MoveToHueAndSaturation{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveToColorId, &MoveToColor{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveColorId, &MoveColor{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, StepColorId, &StepColor{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveToColorTemperatureId, &MoveToColorTemperature{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, EnhancedMoveToHueId, &EnhancedMoveToHue{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, EnhancedMoveHueId, &EnhancedMoveHue{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, EnhancedStepHueId, &EnhancedStepHue{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, EnhancedMoveToHueAndSaturationId, &EnhancedMoveToHueAndSaturation{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, ColorLoopSetId, &ColorLoopSet{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveColorTemperatureId, &MoveColorTemperature{})
	cr.RegisterLocal(zcl.ColorControlId, zigbee.NoManufacturer, zcl.ClientToServer, StepColorTemperatureId, &StepColorTemperature{})
}
