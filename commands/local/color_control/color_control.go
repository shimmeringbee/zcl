package color_control

import (
	"github.com/shimmeringbee/zcl"
)

const (
	CurrentHue             = zcl.AttributeID(0x0000)
	CurrentSaturation      = zcl.AttributeID(0x0001)
	RemainingTime          = zcl.AttributeID(0x0002)
	CurrentX               = zcl.AttributeID(0x0003)
	CurrentY               = zcl.AttributeID(0x0004)
	DriftCompensation      = zcl.AttributeID(0x0005)
	CompensationText       = zcl.AttributeID(0x0006)
	ColorTemperatureMireds = zcl.AttributeID(0x0007)
	ColorMode              = zcl.AttributeID(0x0008)

	EnhancedCurrentHue         = zcl.AttributeID(0x4000)
	EnhancedColorMode          = zcl.AttributeID(0x4001)
	ColorLoopActive            = zcl.AttributeID(0x4002)
	ColorLoopDirection         = zcl.AttributeID(0x4003)
	ColorLoopTime              = zcl.AttributeID(0x4004)
	ColorLoopStartEnhancedHue  = zcl.AttributeID(0x4005)
	ColorLoopStoredEnhancedHue = zcl.AttributeID(0x4006)
	ColorCapabilities          = zcl.AttributeID(0x400a)
	ColorTempPhysicalMinMireds = zcl.AttributeID(0x400b)
	ColorTempPhysicalMaxMireds = zcl.AttributeID(0x400c)

	NumberOfPrimaries = zcl.AttributeID(0x0010)
	Primary1X         = zcl.AttributeID(0x0011)
	Primary1Y         = zcl.AttributeID(0x0012)
	Primary1Intensity = zcl.AttributeID(0x0013)
	Primary2X         = zcl.AttributeID(0x0015)
	Primary2Y         = zcl.AttributeID(0x0016)
	Primary2Intensity = zcl.AttributeID(0x0017)
	Primary3X         = zcl.AttributeID(0x0019)
	Primary3Y         = zcl.AttributeID(0x001a)
	Primary3Intensity = zcl.AttributeID(0x001b)

	Primary4X         = zcl.AttributeID(0x0020)
	Primary4Y         = zcl.AttributeID(0x0021)
	Primary4Intensity = zcl.AttributeID(0x0022)
	Primary5X         = zcl.AttributeID(0x0024)
	Primary5Y         = zcl.AttributeID(0x0025)
	Primary5Intensity = zcl.AttributeID(0x0026)
	Primary6X         = zcl.AttributeID(0x0028)
	Primary6Y         = zcl.AttributeID(0x0029)
	Primary6Intensity = zcl.AttributeID(0x002a)
)

const (
	MoveToHueId                      = zcl.CommandIdentifier(0x00)
	MoveHueId                        = zcl.CommandIdentifier(0x01)
	StepHueId                        = zcl.CommandIdentifier(0x02)
	MoveToSaturationId               = zcl.CommandIdentifier(0x03)
	MoveSaturationId                 = zcl.CommandIdentifier(0x04)
	StepSaturationId                 = zcl.CommandIdentifier(0x05)
	MoveToHueAndSaturationId         = zcl.CommandIdentifier(0x06)
	MoveToColorId                    = zcl.CommandIdentifier(0x07)
	MoveColorId                      = zcl.CommandIdentifier(0x08)
	StepColorId                      = zcl.CommandIdentifier(0x09)
	MoveToColorTemperatureId         = zcl.CommandIdentifier(0x0a)
	EnhancedMoveToHueId              = zcl.CommandIdentifier(0x40)
	EnhancedMoveHueId                = zcl.CommandIdentifier(0x41)
	EnhancedStepHueId                = zcl.CommandIdentifier(0x42)
	EnhancedMoveToHueAndSaturationId = zcl.CommandIdentifier(0x43)
	ColorLoopSetId                   = zcl.CommandIdentifier(0x44)
	StopMoveStepId                   = zcl.CommandIdentifier(0x47)
	MoveColorTemperatureId           = zcl.CommandIdentifier(0x4b)
	StepColorTemperatureId           = zcl.CommandIdentifier(0x4c)
)

type Direction uint8

const (
	DirectionShortestDistance Direction = 0x00
	DirectionLongestDistance  Direction = 0x01
	DirectionUp               Direction = 0x02
	DirectionDown             Direction = 0x03
)

type MoveToHue struct {
	Hue            uint8
	Direction      Direction
	TransitionTime uint16
}

type MoveMode uint8

const (
	MoveModeStop MoveMode = 0x00
	MoveModeUp   MoveMode = 0x01
	MoveModeDown MoveMode = 0x03
)

type MoveHue struct {
	MoveMode MoveMode
	Rate     uint8
}

type StepMode uint8

const (
	StepModeUp   StepMode = 0x01
	StepModeDown StepMode = 0x03
)

type StepHue struct {
	StepMode       StepMode
	StepSize       uint8
	TransitionTime uint8
}

type MoveToSaturation struct {
	Saturation     uint8
	TransitionTime uint16
}

type MoveSaturation struct {
	MoveMode MoveMode
	Rate     uint8
}

type StepSaturation struct {
	StepMode       StepMode
	StepSize       uint8
	TransitionTime uint8
}

type MoveToHueAndSaturation struct {
	Hue            uint8
	Saturation     uint8
	TransitionTime uint16
}

type MoveToColor struct {
	ColorX         uint16
	ColorY         uint16
	TransitionTime uint16
}

type MoveColor struct {
	RateX int16
	RateY int16
}

type StepColor struct {
	StepX          int16
	StepY          int16
	TransitionTime uint16
}

type MoveToColorTemperature struct {
	ColorTemperatureMireds uint16
	TransitionTime         uint16
}

type EnhancedMoveToHue struct {
	EnhancedHue    uint16
	Direction      Direction
	TransitionTime uint16
}

type EnhancedMoveHue struct {
	MoveMode MoveMode
	Rate     uint16
}

type EnhancedStepHue struct {
	StepMode       StepMode
	StepSize       uint16
	TransitionTime uint16
}

type EnhancedMoveToHueAndSaturation struct {
	EnhancedHue    uint16
	Saturation     uint8
	TransitionTime uint16
}

type ColorLoopAction uint8

const (
	DeactivateColorLoop                   ColorLoopAction = 0x00
	ActivateFromColorLoopStartEnhancedHue ColorLoopAction = 0x01
	ActivateFromEnhancedCurrentHue        ColorLoopAction = 0x02
)

type ColorLoopDirectionField uint8

const (
	Decrement ColorLoopDirectionField = 0x00
	Increment ColorLoopDirectionField = 0x01
)

type ColorLoopSet struct {
	Reserved        uint8 `bcfieldwidth:"4"`
	UpdateStartHue  bool  `bcfieldwidth:"1"`
	UpdateTime      bool  `bcfieldwidth:"1"`
	UpdateDirection bool  `bcfieldwidth:"1"`
	UpdateAction    bool  `bcfieldwidth:"1"`
	Action          ColorLoopAction
	Direction       ColorLoopDirectionField
	Time            uint16
	StartHue        uint16
}

type MoveColorTemperature struct {
	MoveMode                      MoveMode
	Rate                          uint16
	ColorTemperatureMinimumMireds uint16
	ColorTemperatureMaximumMireds uint16
}

type StepColorTemperature struct {
	StepMode                      StepMode
	StepSize                      uint16
	TransitionTime                uint16
	ColorTemperatureMinimumMireds uint16
	ColorTemperatureMaximumMireds uint16
}
