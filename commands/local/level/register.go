package level

import (
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
)

func Register(cr *zcl.CommandRegistry) {
	cr.RegisterLocal(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveToLevelId, &MoveToLevel{})
	cr.RegisterLocal(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveId, &Move{})
	cr.RegisterLocal(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, StepId, &Step{})
	cr.RegisterLocal(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, StopId, &Stop{})
	cr.RegisterLocal(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveToLevelWithOnOffId, &MoveToLevelWithOnOff{})
	cr.RegisterLocal(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, MoveWithOnOffId, &MoveWithOnOff{})
	cr.RegisterLocal(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, StepWithOnOffId, &StepWithOnOff{})
	cr.RegisterLocal(zcl.LevelControlId, zigbee.NoManufacturer, zcl.ClientToServer, StopWithOnOffId, &StopWithOnOff{})
}
