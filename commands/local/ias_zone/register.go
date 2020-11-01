package ias_zone

import (
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
)

func Register(cr *zcl.CommandRegistry) {
	cr.RegisterLocal(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ClientToServer, ZoneEnrollResponseId, &ZoneEnrollResponse{})
	cr.RegisterLocal(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ClientToServer, InitiateNormalOperationModeId, &InitiateNormalOperationMode{})
	cr.RegisterLocal(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ClientToServer, InitiateTestModeId, &InitiateTestMode{})

	cr.RegisterLocal(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ServerToClient, ZoneStatusChangeNotificationId, &ZoneStatusChangeNotification{})
	cr.RegisterLocal(zcl.IASZoneId, zigbee.NoManufacturer, zcl.ServerToClient, ZoneEnrollRequestId, &ZoneEnrollRequest{})
}
