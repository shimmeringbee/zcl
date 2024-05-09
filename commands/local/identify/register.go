package identify

import (
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
)

func Register(cr *zcl.CommandRegistry) {
	cr.RegisterLocal(zcl.IdentifyId, zigbee.NoManufacturer, zcl.ClientToServer, IdentifyId, &Identify{})
	cr.RegisterLocal(zcl.IdentifyId, zigbee.NoManufacturer, zcl.ClientToServer, IdentifyQueryId, &IdentifyQuery{})
	cr.RegisterLocal(zcl.IdentifyId, zigbee.NoManufacturer, zcl.ClientToServer, TriggerEffectId, &TriggerEffect{})

	cr.RegisterLocal(zcl.IdentifyId, zigbee.NoManufacturer, zcl.ServerToClient, IdentifyQueryResponseId, &IdentifyQueryResponse{})
}
