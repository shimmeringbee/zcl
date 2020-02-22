package zcl

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

type ReadAttributesStructuredSelector struct {
	Reserved uint8    `bcfieldwidth:"4"`
	Index    []uint16 `bcsliceprefix:"4"`
}

type ReadAttributesStructuredRecord struct {
	Identifier AttributeIdentifier
	Selector   ReadAttributesStructuredSelector
}

type ReadAttributesStructured struct {
	Records []ReadAttributesStructuredRecord
}
