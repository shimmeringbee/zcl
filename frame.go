package zcl

import "github.com/shimmeringbee/zigbee"

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
	ReadAttributesID                     CommandIdentifier = 0x00
	ReadAttributesResponseID             CommandIdentifier = 0x01
	WriteAttributesID                    CommandIdentifier = 0x02
	WriteAttributesUndividedID           CommandIdentifier = 0x03
	WriteAttributesResponseID            CommandIdentifier = 0x04
	WriteAttributesNoResponseID          CommandIdentifier = 0x05
	ConfigureReportingID                 CommandIdentifier = 0x06
	ConfigureReportingResponseID         CommandIdentifier = 0x07
	ReadReportingConfigurationID         CommandIdentifier = 0x08
	ReadReportingConfigurationResponseID CommandIdentifier = 0x09
	ReportAttributesID                   CommandIdentifier = 0x0a
	DefaultResponseID                    CommandIdentifier = 0x0b
	DiscoverAttributesID                 CommandIdentifier = 0x0c
	DiscoverAttributesResponseID         CommandIdentifier = 0x0d
	ReadAttributesStructuredID           CommandIdentifier = 0x0e
	WriteAttributesStructuredID          CommandIdentifier = 0x0f
	WriteAttributesStructuredResponseID  CommandIdentifier = 0x10
	DiscoverCommandsReceivedID           CommandIdentifier = 0x11
	DiscoverCommandsReceivedResponseID   CommandIdentifier = 0x12
	DiscoverCommandsGeneratedID          CommandIdentifier = 0x13
	DiscoverCommandsGeneratedResponseID  CommandIdentifier = 0x14
	DiscoverAttributesExtendedID         CommandIdentifier = 0x15
	DiscoverAttributesExtendedResponseID CommandIdentifier = 0x16
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

type ZCLMessage struct {
	FrameType           FrameType
	Direction           Direction
	TransactionSequence uint8
	Manufacturer        uint16
	ClusterID           zigbee.ClusterID
	SourceEndpoint      zigbee.Endpoint
	DestinationEndpoint zigbee.Endpoint
	Command             interface{}
}

func (z ZCLMessage) isManufacturerSpecific() bool {
	return z.Manufacturer > 0
}
