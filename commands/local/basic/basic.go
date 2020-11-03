package basic

import "github.com/shimmeringbee/zcl"

const (
	ZCLVersion          = zcl.AttributeID(0x0000)
	ApplicationVersion  = zcl.AttributeID(0x0001)
	StackVersion        = zcl.AttributeID(0x0002)
	HWVersion           = zcl.AttributeID(0x0003)
	ManufacturerName    = zcl.AttributeID(0x0004)
	ModelIdentifier     = zcl.AttributeID(0x0005)
	DateCode            = zcl.AttributeID(0x0006)
	PowerSource         = zcl.AttributeID(0x0007)
	LocationDescription = zcl.AttributeID(0x0010)
	PhysicalEnvironment = zcl.AttributeID(0x0011)
	DeviceEnabled       = zcl.AttributeID(0x0012)
	AlarmMask           = zcl.AttributeID(0x0013)
	DisableLocalConfig  = zcl.AttributeID(0x0014)
	SWBuildID           = zcl.AttributeID(0x4000)
)
