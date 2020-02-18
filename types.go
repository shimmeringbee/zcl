package zcl

/*
 * Zigbee Cluster List data types, as per 2.6.2 in ZCL Revision 6 (14 January 2016).
 * Downloaded From: https://zigbeealliance.org/developer_resources/zigbee-cluster-library/
 */

const (
	TypeNull AttributeDataType = 0x00

	TypeData8  AttributeDataType = 0x08
	TypeData16 AttributeDataType = 0x09
	TypeData24 AttributeDataType = 0x0a
	TypeData32 AttributeDataType = 0x0b
	TypeData40 AttributeDataType = 0x0c
	TypeData48 AttributeDataType = 0x0d
	TypeData56 AttributeDataType = 0x0e
	TypeData64 AttributeDataType = 0x0f

	TypeBoolean AttributeDataType = 0x10

	TypeBitmap8  AttributeDataType = 0x18
	TypeBitmap16 AttributeDataType = 0x19
	TypeBitmap24 AttributeDataType = 0x1a
	TypeBitmap32 AttributeDataType = 0x1b
	TypeBitmap40 AttributeDataType = 0x1c
	TypeBitmap48 AttributeDataType = 0x1d
	TypeBitmap56 AttributeDataType = 0x1e
	TypeBitmap64 AttributeDataType = 0x1f

	TypeUnsignedInt8  AttributeDataType = 0x20
	TypeUnsignedInt16 AttributeDataType = 0x21
	TypeUnsignedInt24 AttributeDataType = 0x22
	TypeUnsignedInt32 AttributeDataType = 0x23
	TypeUnsignedInt40 AttributeDataType = 0x24
	TypeUnsignedInt48 AttributeDataType = 0x25
	TypeUnsignedInt56 AttributeDataType = 0x26
	TypeUnsignedInt64 AttributeDataType = 0x27

	TypeSignedInt8  AttributeDataType = 0x28
	TypeSignedInt16 AttributeDataType = 0x29
	TypeSignedInt24 AttributeDataType = 0x2a
	TypeSignedInt32 AttributeDataType = 0x2b
	TypeSignedInt40 AttributeDataType = 0x2c
	TypeSignedInt48 AttributeDataType = 0x2d
	TypeSignedInt56 AttributeDataType = 0x2e
	TypeSignedInt64 AttributeDataType = 0x2f

	TypeEnum8  AttributeDataType = 0x30
	TypeEnum16 AttributeDataType = 0x31

	TypeFloatSemi   AttributeDataType = 0x38
	TypeFloatSingle AttributeDataType = 0x39
	TypeFloatDouble AttributeDataType = 0x3a

	TypeStringOctet8      AttributeDataType = 0x41
	TypeStringCharacter8  AttributeDataType = 0x42
	TypeStringOctet16     AttributeDataType = 0x43
	TypeStringCharacter16 AttributeDataType = 0x44

	TypeArray     AttributeDataType = 0x48
	TypeStructure AttributeDataType = 0x4c
	TypeSet       AttributeDataType = 0x50
	TypeBag       AttributeDataType = 0x51

	TypeTimeOfDay AttributeDataType = 0xe0
	TypeDate      AttributeDataType = 0xe1
	TypeUTCTime   AttributeDataType = 0xe2

	TypeClusterID   AttributeDataType = 0xe9
	TypeAttributeID AttributeDataType = 0xea
	TypeBACnetOID   AttributeDataType = 0xeb

	TypeIEEEAddress    AttributeDataType = 0xf0
	TypeSecurityKey128 AttributeDataType = 0xf1
	TypeUnknown        AttributeDataType = 0xff
)
