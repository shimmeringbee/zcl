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
