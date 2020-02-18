package zcl

type ReadAttributes struct {
	Identifier []AttributeIdentifier
}

type ReadAttributeResponseStatusRecord struct {
	Identifier    AttributeIdentifier
	Status        uint8
	DataTypeValue *AttributeDataTypeValue `bcincludeif:"Status==0"`
}

type ReadAttributesResponse struct {
	Records []ReadAttributeResponseStatusRecord
}
