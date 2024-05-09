package identify

import "github.com/shimmeringbee/zcl"

const (
	IdentifyTime = zcl.AttributeID(0x0000)
)

const (
	IdentifyId      = zcl.CommandIdentifier(0x00)
	IdentifyQueryId = zcl.CommandIdentifier(0x01)
	TriggerEffectId = zcl.CommandIdentifier(0x02)

	IdentifyQueryResponseId = zcl.CommandIdentifier(0x00)
)

type Identify struct {
	IdentifyTime uint16
}

type IdentifyQuery struct{}

type TriggerEffect struct {
	EffectIdentifier uint8
	EffectVariant    uint8
}

type IdentifyQueryResponse struct {
	Timeout uint16
}
