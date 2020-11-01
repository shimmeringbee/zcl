package ias_zone

import (
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
)

const (
	ZoneState                              = zcl.AttributeID(0x0000)
	ZoneType                               = zcl.AttributeID(0x0001)
	ZoneStatus                             = zcl.AttributeID(0x0002)
	IASCIEAddress                          = zcl.AttributeID(0x0010)
	ZoneID                                 = zcl.AttributeID(0x0011)
	NumberOfZoneSensitivityLevelsSupported = zcl.AttributeID(0x0012)
	CurrentZoneSensitivityLevel            = zcl.AttributeID(0x0013)
)

const (
	ZoneEnrollResponseId          = zcl.CommandIdentifier(0x00)
	InitiateNormalOperationModeId = zcl.CommandIdentifier(0x01)
	InitiateTestModeId            = zcl.CommandIdentifier(0x02)

	ZoneStatusChangeNotificationId = zcl.CommandIdentifier(0x00)
	ZoneEnrollRequestId            = zcl.CommandIdentifier(0x01)
)

type ZoneEnrollResponse struct {
	ResponseCode uint8
	ZoneID       uint8
}

type InitiateNormalOperationMode struct{}

type InitiateTestMode struct {
	TestModeDuration            uint8
	CurrentZoneSensitivityLevel uint8
}

type ZoneStatusChangeNotification struct {
	Reserved           uint8 `bcfieldwidth:"6"`
	BatteryDefect      bool  `bcfieldwidth:"1"`
	TestMode           bool  `bcfieldwidth:"1"`
	ACMainsFault       bool  `bcfieldwidth:"1"`
	Trouble            bool  `bcfieldwidth:"1"`
	RestoreReports     bool  `bcfieldwidth:"1"`
	SupervisionReports bool  `bcfieldwidth:"1"`
	BatteryLow         bool  `bcfieldwidth:"1"`
	Tamper             bool  `bcfieldwidth:"1"`
	Alarm2             bool  `bcfieldwidth:"1"`
	Alarm1             bool  `bcfieldwidth:"1"`
	ExtendedStatus     uint8
	ZoneID             uint8
	Delay              uint16
}

type ZoneEnrollRequest struct {
	ZoneType         uint16
	ManufacturerCode zigbee.ManufacturerCode
}
