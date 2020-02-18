package zcl

/*
 * Zigbee Cluster List data types, as per 2.6.2 in ZCL Revision 6 (14 January 2016).
 * Downloaded From: https://zigbeealliance.org/developer_resources/zigbee-cluster-library/
 */

const (
	TypeNull byte = 0x00

	TypeData8  byte = 0x08
	TypeData16 byte = 0x09
	TypeData24 byte = 0x0a
	TypeData32 byte = 0x0b
	TypeData40 byte = 0x0c
	TypeData48 byte = 0x0d
	TypeData56 byte = 0x0e
	TypeData64 byte = 0x0f

	TypeBoolean byte = 0x10

	TypeBitmap8  byte = 0x18
	TypeBitmap16 byte = 0x19
	TypeBitmap24 byte = 0x1a
	TypeBitmap32 byte = 0x1b
	TypeBitmap40 byte = 0x1c
	TypeBitmap48 byte = 0x1d
	TypeBitmap56 byte = 0x1e
	TypeBitmap64 byte = 0x1f

	TypeUnsignedInt8  byte = 0x20
	TypeUnsignedInt16 byte = 0x21
	TypeUnsignedInt24 byte = 0x22
	TypeUnsignedInt32 byte = 0x23
	TypeUnsignedInt40 byte = 0x24
	TypeUnsignedInt48 byte = 0x25
	TypeUnsignedInt56 byte = 0x26
	TypeUnsignedInt64 byte = 0x27

	TypeSignedInt8  byte = 0x28
	TypeSignedInt16 byte = 0x29
	TypeSignedInt24 byte = 0x2a
	TypeSignedInt32 byte = 0x2b
	TypeSignedInt40 byte = 0x2c
	TypeSignedInt48 byte = 0x2d
	TypeSignedInt56 byte = 0x2e
	TypeSignedInt64 byte = 0x2f

	TypeEnum8  byte = 0x30
	TypeEnum16 byte = 0x31

	TypeFloatSemi   byte = 0x38
	TypeFloatSingle byte = 0x39
	TypeFloatDouble byte = 0x3a

	TypeStringOctet8      byte = 0x41
	TypeStringCharacter8  byte = 0x42
	TypeStringOctet16     byte = 0x43
	TypeStringCharacter16 byte = 0x44

	TypeArray     byte = 0x48
	TypeStructure byte = 0x4c
	TypeSet       byte = 0x50
	TypeBag       byte = 0x51

	TypeTimeOfDay byte = 0xe0
	TypeDate      byte = 0xe1
	TypeUTCTime   byte = 0xe2

	TypeClusterID   byte = 0xe9
	TypeAttributeID byte = 0xea
	TypeBACnetOID   byte = 0xeb

	TypeIEEEAddress    byte = 0xf0
	TypeSecurityKey128 byte = 0xf1
	TypeUnknown        byte = 0xff
)
