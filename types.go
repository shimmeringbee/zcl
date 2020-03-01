package zcl

import (
	"errors"
	"fmt"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
)

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

var DiscreteTypes = map[AttributeDataType]bool{
	TypeNull: false,

	TypeData8:  true,
	TypeData16: true,
	TypeData24: true,
	TypeData32: true,
	TypeData40: true,
	TypeData48: true,
	TypeData56: true,
	TypeData64: true,

	TypeBoolean: true,

	TypeBitmap8:  true,
	TypeBitmap16: true,
	TypeBitmap24: true,
	TypeBitmap32: true,
	TypeBitmap40: true,
	TypeBitmap48: true,
	TypeBitmap56: true,
	TypeBitmap64: true,

	TypeUnsignedInt8:  false,
	TypeUnsignedInt16: false,
	TypeUnsignedInt24: false,
	TypeUnsignedInt32: false,
	TypeUnsignedInt40: false,
	TypeUnsignedInt48: false,
	TypeUnsignedInt56: false,
	TypeUnsignedInt64: false,

	TypeSignedInt8:  false,
	TypeSignedInt16: false,
	TypeSignedInt24: false,
	TypeSignedInt32: false,
	TypeSignedInt40: false,
	TypeSignedInt48: false,
	TypeSignedInt56: false,
	TypeSignedInt64: false,

	TypeEnum8:  true,
	TypeEnum16: true,

	TypeFloatSemi:   false,
	TypeFloatSingle: false,
	TypeFloatDouble: false,

	TypeStringOctet8:      true,
	TypeStringCharacter8:  true,
	TypeStringOctet16:     true,
	TypeStringCharacter16: true,

	TypeArray:     true,
	TypeStructure: true,
	TypeSet:       true,
	TypeBag:       true,

	TypeTimeOfDay: false,
	TypeDate:      false,
	TypeUTCTime:   false,

	TypeClusterID:   true,
	TypeAttributeID: true,
	TypeBACnetOID:   true,

	TypeIEEEAddress:    true,
	TypeSecurityKey128: true,
	TypeUnknown:        false,
}

type AttributeDataType byte
type AttributeIdentifier uint16

type AttributeDataTypeValue struct {
	DataType AttributeDataType
	Value    interface{}
}

func (a *AttributeDataTypeValue) Marshal(bb *bitbuffer.BitBuffer) error {
	if err := bb.WriteByte(byte(a.DataType)); err != nil {
		return err
	}

	switch a.DataType {
	case TypeNull:
		return nil
	case TypeData8:
		return a.marshalData(bb, 1)
	case TypeData16:
		return a.marshalData(bb, 2)
	case TypeData24:
		return a.marshalData(bb, 3)
	case TypeData32:
		return a.marshalData(bb, 4)
	case TypeData40:
		return a.marshalData(bb, 5)
	case TypeData48:
		return a.marshalData(bb, 6)
	case TypeData56:
		return a.marshalData(bb, 7)
	case TypeData64:
		return a.marshalData(bb, 8)
	case TypeBoolean:
		return a.marshalBoolean(bb)
	default:
		return fmt.Errorf("unsupported ZCL type to marshal: %d", a.DataType)
	}
}

func (a *AttributeDataTypeValue) marshalData(bb *bitbuffer.BitBuffer, size int) error {
	data, ok := a.Value.([]byte)

	if !ok {
		return errors.New("could not cast value")
	}

	if len(data) != size {
		return fmt.Errorf("data array provided does not match output size")
	}

	for i := size - 1; i >= 0; i-- {
		if err := bb.WriteByte(data[i]); err != nil {
			return err
		}
	}

	return nil
}

func (a *AttributeDataTypeValue) marshalBoolean(bb *bitbuffer.BitBuffer) error {
	data, ok := a.Value.(bool)

	if !ok {
		return errors.New("could not cast value")
	}

	if data {
		return bb.WriteByte(0x01)
	} else {
		return bb.WriteByte(0x00)
	}
}

func (a *AttributeDataTypeValue) Unmarshal(bb *bitbuffer.BitBuffer) error {
	if dt, err := bb.ReadByte(); err != nil {
		return err
	} else {
		a.DataType = AttributeDataType(dt)
	}

	switch a.DataType {
	case TypeNull:
		return nil
	case TypeData8:
		return a.unmarshalData(bb, 1)
	case TypeData16:
		return a.unmarshalData(bb, 2)
	case TypeData24:
		return a.unmarshalData(bb, 3)
	case TypeData32:
		return a.unmarshalData(bb, 4)
	case TypeData40:
		return a.unmarshalData(bb, 5)
	case TypeData48:
		return a.unmarshalData(bb, 6)
	case TypeData56:
		return a.unmarshalData(bb, 7)
	case TypeData64:
		return a.unmarshalData(bb, 8)
	case TypeBoolean:
		return a.unmarshalBoolean(bb)
	default:
		return fmt.Errorf("unsupported ZCL type to unmarshal: %d", a.DataType)
	}
}

func (a *AttributeDataTypeValue) unmarshalData(bb *bitbuffer.BitBuffer, size int) error {
	data := make([]byte, size)

	for i := size - 1; i >= 0; i-- {
		if b, err := bb.ReadByte(); err != nil {
			return err
		} else {
			data[i] = b
		}
	}

	a.Value = data

	return nil
}

func (a *AttributeDataTypeValue) unmarshalBoolean(bb *bitbuffer.BitBuffer) error {
	if data, err := bb.ReadByte(); err != nil {
		return err
	} else {
		a.Value = data != 0x00
		return nil
	}
}
