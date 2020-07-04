package onoff

import (
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
)

func Register(cr *zcl.CommandRegistry) {
	cr.RegisterLocal(zcl.OnOffId, zigbee.NoManufacturer, OffId, &Off{})
	cr.RegisterLocal(zcl.OnOffId, zigbee.NoManufacturer, OnId, &On{})
	cr.RegisterLocal(zcl.OnOffId, zigbee.NoManufacturer, ToggleId, &Toggle{})
	cr.RegisterLocal(zcl.OnOffId, zigbee.NoManufacturer, OffWithEffectId, &OffWithEffect{})
	cr.RegisterLocal(zcl.OnOffId, zigbee.NoManufacturer, OnWithRecallGlobalSceneId, &OnWithRecallGlobalScene{})
	cr.RegisterLocal(zcl.OnOffId, zigbee.NoManufacturer, OnWithTimedOffId, &OnWithTimedOff{})
}
