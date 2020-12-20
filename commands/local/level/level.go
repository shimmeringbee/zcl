package level

import "github.com/shimmeringbee/zcl"

const (
	CurrentLevel        = zcl.AttributeID(0x0000)
	RemainingTime       = zcl.AttributeID(0x0001)
	OnOffTransitionTime = zcl.AttributeID(0x0010)
	OnLevel             = zcl.AttributeID(0x0011)
	OnTransitionTime    = zcl.AttributeID(0x0012)
	OffTransitionTime   = zcl.AttributeID(0x0013)
	DefaultMoveRate     = zcl.AttributeID(0x0014)
)

const (
	MoveToLevelId          = zcl.CommandIdentifier(0x00)
	MoveId                 = zcl.CommandIdentifier(0x01)
	StepId                 = zcl.CommandIdentifier(0x02)
	StopId                 = zcl.CommandIdentifier(0x03)
	MoveToLevelWithOnOffId = zcl.CommandIdentifier(0x04)
	MoveWithOnOffId        = zcl.CommandIdentifier(0x05)
	StepWithOnOffId        = zcl.CommandIdentifier(0x06)
	StopWithOnOffId        = zcl.CommandIdentifier(0x07)
)

type MoveToLevel struct {
	Level          uint8
	TransitionTime uint16
}

type Move struct {
	MoveMode uint8
	Rate     uint8
}

type Step struct {
	StepMode       uint8
	StepSize       uint8
	TransitionTime uint16
}

type Stop struct{}

type MoveToLevelWithOnOff struct {
	Level          uint8
	TransitionTime uint16
}

type MoveWithOnOff struct {
	MoveMode uint8
	Rate     uint8
}

type StepWithOnOff struct {
	StepMode       uint8
	StepSize       uint8
	TransitionTime uint16
}

type StopWithOnOff struct{}
