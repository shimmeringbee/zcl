package zcl

type FrameType uint8

const (
	FrameGlobal FrameType = 0x00
	FrameLocal  FrameType = 0x01
)

type Direction uint8

const (
	ClientToServer Direction = 0
	ServerToClient Direction = 1
)

type CommandIdentifier uint8

const (
	ReadAttributes                     CommandIdentifier = 0x00
	ReadAttributesResponse             CommandIdentifier = 0x01
	WriteAttributes                    CommandIdentifier = 0x02
	WriteAttributesUndivided           CommandIdentifier = 0x03
	WriteAttributesResponse            CommandIdentifier = 0x04
	WriteAttributesNoResponse          CommandIdentifier = 0x05
	ConfigureReporting                 CommandIdentifier = 0x06
	ConfigureReportingResponse         CommandIdentifier = 0x07
	ReadReportingConfiguration         CommandIdentifier = 0x08
	ReadReportingConfigurationResponse CommandIdentifier = 0x09
	ReportAttributes                   CommandIdentifier = 0x0a
	DefaultResponse                    CommandIdentifier = 0x0b
	DiscoverAttributes                 CommandIdentifier = 0x0c
	DiscoverAttributesResponse         CommandIdentifier = 0x0d
	ReadAttributesStructured           CommandIdentifier = 0x0e
	WriteAttributesStructured          CommandIdentifier = 0x0f
	WriteAttributesStructuredResponse  CommandIdentifier = 0x10
	DiscoverCommandsReceived           CommandIdentifier = 0x11
	DiscoverCommandsReceivedResponse   CommandIdentifier = 0x12
	DiscoverCommandsGenerated          CommandIdentifier = 0x13
	DiscoverCommandsGeneratedResponse  CommandIdentifier = 0x14
	DiscoverAttributesExtended         CommandIdentifier = 0x15
	DiscoverAttributesExtendedResponse CommandIdentifier = 0x16
)

type Control struct {
	Reserved               uint8     `bcfieldwidth:"3"`
	DisableDefaultResponse bool      `bcfieldwidth:"1"`
	Direction              Direction `bcfieldwidth:"1"`
	ManufacturerSpecific   bool      `bcfieldwidth:"1"`
	FrameType              FrameType `bcfieldwidth:"2"`
}

type Header struct {
	Control             Control
	Manufacturer        uint16 `bcincludeif:".Control.ManufacturerSpecific"`
	TransactionSequence uint8
	CommandIdentifier   CommandIdentifier
}
