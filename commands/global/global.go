package global

import . "github.com/shimmeringbee/zcl"

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

type ReadAttributes struct {
	Identifier []AttributeIdentifier
}

type ReadAttributeResponseRecord struct {
	Identifier    AttributeIdentifier
	Status        uint8
	DataTypeValue *AttributeDataTypeValue `bcincludeif:"Status==0"`
}

type ReadAttributesResponse struct {
	Records []ReadAttributeResponseRecord
}

type WriteAttributes struct {
	Records []WriteAttributesRecord
}

type WriteAttributesRecord struct {
	Identifier    AttributeIdentifier
	DataTypeValue *AttributeDataTypeValue
}

type WriteAttributesResponseRecord struct {
	Status     uint8
	Identifier AttributeIdentifier
}

type WriteAttributesResponse struct {
	Records []WriteAttributesResponseRecord
}

type WriteAttributesUndivided WriteAttributes

type WriteAttributesNoResponse WriteAttributes

type ConfigureReportingRecord struct {
	Direction        uint8
	Identifier       AttributeIdentifier
	DataType         AttributeDataType `bcincludeif:"Direction==0"`
	MinimumInterval  uint16            `bcincludeif:"Direction==0"`
	MaximumInterval  uint16            `bcincludeif:"Direction==0"`
	ReportableChange interface{}       `bcincludeif:"Direction==0"`
	Timeout          uint16            `bcincludeif:"Direction==1"`
}

type ConfigureReporting struct {
	Records []ConfigureReportingRecord
}

type ConfigureReportingResponseRecord struct {
	Status     uint8
	Direction  uint8
	Identifier AttributeIdentifier
}

type ConfigureReportingResponse struct {
	Records []ConfigureReportingResponseRecord
}

type ReadReportingConfigurationRecord struct {
	Direction  uint8
	Identifier AttributeIdentifier
}

type ReadReportingConfiguration struct {
	Records []ReadReportingConfigurationRecord
}

type ReadReportingConfigurationResponseRecord struct {
	Status           uint8
	Direction        uint8
	Identifier       AttributeIdentifier
	DataType         AttributeDataType `bcincludeif:"Direction==0"`
	MinimumInterval  uint16            `bcincludeif:"Direction==0"`
	MaximumInterval  uint16            `bcincludeif:"Direction==0"`
	ReportableChange interface{}       `bcincludeif:"Direction==0"`
	Timeout          uint16            `bcincludeif:"Direction==1"`
}

type ReadReportingConfigurationResponse struct {
	Records []ReadReportingConfigurationResponseRecord
}

type ReportAttributesRecord struct {
	Identifier    AttributeIdentifier
	DataTypeValue *AttributeDataTypeValue
}

type ReportAttributes struct {
	Records []ReportAttributesRecord
}

type DefaultResponse struct {
	CommandIdentifier uint8
	Status            uint8
}

type DiscoverAttributes struct {
	StartAttributeIdentifier  uint16
	MaximumNumberOfAttributes uint8
}

type DiscoverAttributesResponseRecord struct {
	Identifier AttributeIdentifier
	DataType   AttributeDataType
}

type DiscoverAttributesResponse struct {
	DiscoveryComplete bool
	Records           []DiscoverAttributesResponseRecord
}

type ReadAttributesStructuredRecord struct {
	Identifier AttributeIdentifier
	Selector   Selector
}

type ReadAttributesStructured struct {
	Records []ReadAttributesStructuredRecord
}

type Selector struct {
	BagSetOperation uint8    `bcfieldwidth:"4"`
	Index           []uint16 `bcsliceprefix:"4"`
}

type WriteAttributesStructuredRecord struct {
	Identifier    AttributeIdentifier
	Selector      Selector
	DataTypeValue *AttributeDataTypeValue
}

type WriteAttributesStructured struct {
	Records []WriteAttributesStructuredRecord
}

type WriteAttributesStructuredResponseRecord struct {
	Status     uint8
	Identifier AttributeIdentifier
	Selector   Selector
}

type WriteAttributesStructuredResponse struct {
	Records []WriteAttributesStructuredResponseRecord
}

type DiscoverCommandsReceived struct {
	StartCommandIdentifier  uint8
	MaximumNumberOfCommands uint8
}

type DiscoverCommandsReceivedResponse struct {
	DiscoveryComplete bool
	CommandIdentifier []uint8
}

type DiscoverCommandsGenerated struct {
	StartCommandIdentifier  uint8
	MaximumNumberOfCommands uint8
}

type DiscoverCommandsGeneratedResponse struct {
	DiscoveryComplete bool
	CommandIdentifier []uint8
}

type DiscoverAttributesExtended struct {
	StartAttributeIdentifier  uint16
	MaximumNumberOfAttributes uint8
}

type DiscoverAttributesExtendedResponseRecord struct {
	Identifier    AttributeIdentifier
	DataType      AttributeDataType
	AccessControl uint8
}

type DiscoverAttributesExtendedResponse struct {
	DiscoveryComplete bool
	Records           []DiscoverAttributesExtendedResponseRecord
}
