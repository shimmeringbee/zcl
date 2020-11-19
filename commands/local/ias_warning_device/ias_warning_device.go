package ias_warning_device

import "github.com/shimmeringbee/zcl"

const (
	MaxDuration = zcl.AttributeID(0x0000)
)

const (
	StartWarningId = zcl.CommandIdentifier(0x00)
	SquawkId       = zcl.CommandIdentifier(0x01)
)

type WarningMode uint8

const (
	Stop      = WarningMode(0)
	Burglar   = WarningMode(1)
	Fire      = WarningMode(2)
	Emergency = WarningMode(3)
)

type StrobeMode uint8

const (
	NoStrobe          = StrobeMode(0)
	StrobeWithWarning = StrobeMode(1)
)

type StartWarning struct {
	WarningMode     WarningMode `bcfieldwidth:"4"`
	StrobeMode      StrobeMode  `bcfieldwidth:"2"`
	Reserved        uint8       `bcfieldwidth:"2"`
	WarningDuration uint16
}

type SquawkMode uint8

const (
	SystemArmed    = SquawkMode(0)
	SystemDisarmed = SquawkMode(1)
)

type SquawkLevel uint8

const (
	Low      = SquawkLevel(0)
	Medium   = SquawkLevel(1)
	High     = SquawkLevel(2)
	VeryHigh = SquawkLevel(3)
)

type Squawk struct {
	SquawkMode  SquawkMode  `bcfieldwidth:"4"`
	Strobe      bool        `bcfieldwidth:"1"`
	Reserved    uint8       `bcfieldwidth:"1"`
	SquawkLevel SquawkLevel `bcfieldwidth:"2"`
}
