package relative_humidity_measurement

import "github.com/shimmeringbee/zcl"

const (
	MeasuredValue    = zcl.AttributeID(0x0000)
	MinMeasuredValue = zcl.AttributeID(0x0001)
	MaxMeasuredValue = zcl.AttributeID(0x0002)
	Tolerance        = zcl.AttributeID(0x0003)
)
