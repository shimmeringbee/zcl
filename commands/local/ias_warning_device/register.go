package ias_warning_device

import (
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
)

func Register(cr *zcl.CommandRegistry) {
	cr.RegisterLocal(zcl.IASWarningDevicesId, zigbee.NoManufacturer, zcl.ClientToServer, StartWarningId, &StartWarning{})
	cr.RegisterLocal(zcl.IASWarningDevicesId, zigbee.NoManufacturer, zcl.ClientToServer, SquawkId, &Squawk{})
}
