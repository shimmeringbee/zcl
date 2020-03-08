package zcl

import (
	"errors"
	"fmt"
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
	"github.com/shimmeringbee/zigbee"
	"math"
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
type AttributeID uint16

type AttributeDataTypeValue struct {
	DataType AttributeDataType
	Value    interface{}
}

type AttributeSlice struct {
	DataType AttributeDataType
	Values   []interface{}
}

func (a *AttributeDataTypeValue) Marshal(bb *bitbuffer.BitBuffer) error {
	if err := bb.WriteByte(byte(a.DataType)); err != nil {
		return err
	}

	return marshalZCLType(bb, a.DataType, a.Value)
}

func marshalZCLType(bb *bitbuffer.BitBuffer, dt AttributeDataType, v interface{}) error {
	switch dt {
	case TypeNull:
		return nil
	case TypeData8:
		return marshalData(bb, v, 1)
	case TypeData16:
		return marshalData(bb, v, 2)
	case TypeData24:
		return marshalData(bb, v, 3)
	case TypeData32:
		return marshalData(bb, v, 4)
	case TypeData40:
		return marshalData(bb, v, 5)
	case TypeData48:
		return marshalData(bb, v, 6)
	case TypeData56:
		return marshalData(bb, v, 7)
	case TypeData64:
		return marshalData(bb, v, 8)
	case TypeBoolean:
		return marshalBoolean(bb, v)
	case TypeBitmap8:
		return marshalUint(bb, v, 8)
	case TypeBitmap16:
		return marshalUint(bb, v, 16)
	case TypeBitmap24:
		return marshalUint(bb, v, 24)
	case TypeBitmap32:
		return marshalUint(bb, v, 32)
	case TypeBitmap40:
		return marshalUint(bb, v, 40)
	case TypeBitmap48:
		return marshalUint(bb, v, 48)
	case TypeBitmap56:
		return marshalUint(bb, v, 56)
	case TypeBitmap64:
		return marshalUint(bb, v, 64)
	case TypeUnsignedInt8:
		return marshalUint(bb, v, 8)
	case TypeUnsignedInt16:
		return marshalUint(bb, v, 16)
	case TypeUnsignedInt24:
		return marshalUint(bb, v, 24)
	case TypeUnsignedInt32:
		return marshalUint(bb, v, 32)
	case TypeUnsignedInt40:
		return marshalUint(bb, v, 40)
	case TypeUnsignedInt48:
		return marshalUint(bb, v, 48)
	case TypeUnsignedInt56:
		return marshalUint(bb, v, 56)
	case TypeUnsignedInt64:
		return marshalUint(bb, v, 64)
	case TypeSignedInt8:
		return marshalInt(bb, v, 8)
	case TypeSignedInt16:
		return marshalInt(bb, v, 16)
	case TypeSignedInt24:
		return marshalInt(bb, v, 24)
	case TypeSignedInt32:
		return marshalInt(bb, v, 32)
	case TypeSignedInt40:
		return marshalInt(bb, v, 40)
	case TypeSignedInt48:
		return marshalInt(bb, v, 48)
	case TypeSignedInt56:
		return marshalInt(bb, v, 56)
	case TypeSignedInt64:
		return marshalInt(bb, v, 64)
	case TypeEnum8:
		return marshalUint(bb, v, 8)
	case TypeEnum16:
		return marshalUint(bb, v, 16)
	case TypeStringOctet8:
		return marshalString(bb, v, 8)
	case TypeStringOctet16:
		return marshalString(bb, v, 16)
	case TypeStringCharacter8:
		return marshalStringRune(bb, v, 8)
	case TypeStringCharacter16:
		return marshalStringRune(bb, v, 16)
	case TypeTimeOfDay:
		return marshalTimeOfDay(bb, v)
	case TypeDate:
		return marshalDate(bb, v)
	case TypeUTCTime:
		return marshalUTCTime(bb, v)
	case TypeClusterID:
		return marshalClusterID(bb, v)
	case TypeAttributeID:
		return marshalAttributeID(bb, v)
	case TypeIEEEAddress:
		return marshalIEEEAddress(bb, v)
	case TypeSecurityKey128:
		return marshalSecurityKey(bb, v)
	case TypeBACnetOID:
		return marshalBACnetOID(bb, v)
	case TypeStructure:
		return marshalStructure(bb, v)
	case TypeArray, TypeSet, TypeBag:
		return marshalSlice(bb, v)
	case TypeFloatSingle:
		return marshalFloatSingle(bb, v)
	case TypeFloatDouble:
		return marshalFloatDouble(bb, v)
	default:
		return fmt.Errorf("unsupported ZCL type to marshal: %d", dt)
	}
}

func marshalData(bb *bitbuffer.BitBuffer, v interface{}, size int) error {
	data, ok := v.([]byte)

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

func marshalBoolean(bb *bitbuffer.BitBuffer, v interface{}) error {
	data, ok := v.(bool)

	if !ok {
		return errors.New("could not cast value")
	}

	if data {
		return bb.WriteByte(0x01)
	} else {
		return bb.WriteByte(0x00)
	}
}

func marshalUint(bb *bitbuffer.BitBuffer, v interface{}, bitsize int) error {
	switch v := v.(type) {
	case uint:
		return bb.WriteUint(uint64(v), bitbuffer.LittleEndian, bitsize)
	case uint8:
		return bb.WriteUint(uint64(v), bitbuffer.LittleEndian, bitsize)
	case uint16:
		return bb.WriteUint(uint64(v), bitbuffer.LittleEndian, bitsize)
	case uint32:
		return bb.WriteUint(uint64(v), bitbuffer.LittleEndian, bitsize)
	case uint64:
		return bb.WriteUint(v, bitbuffer.LittleEndian, bitsize)
	}

	return errors.New("marshalling uint to ZCL type received unsupported value")
}

func marshalInt(bb *bitbuffer.BitBuffer, v interface{}, bitsize int) error {
	switch v := v.(type) {
	case int:
		return bb.WriteInt(int64(v), bitbuffer.LittleEndian, bitsize)
	case int8:
		return bb.WriteInt(int64(v), bitbuffer.LittleEndian, bitsize)
	case int16:
		return bb.WriteInt(int64(v), bitbuffer.LittleEndian, bitsize)
	case int32:
		return bb.WriteInt(int64(v), bitbuffer.LittleEndian, bitsize)
	case int64:
		return bb.WriteInt(v, bitbuffer.LittleEndian, bitsize)
	}

	return errors.New("marshalling int to ZCL type received unsupported value")
}

func marshalString(bb *bitbuffer.BitBuffer, v interface{}, bitsize int) error {
	data, ok := v.(string)

	if !ok {
		return errors.New("could not cast value")
	}

	return bb.WriteStringLengthPrefixed(data, bitbuffer.LittleEndian, bitsize)
}

func marshalStringRune(bb *bitbuffer.BitBuffer, v interface{}, bitsize int) error {
	data, ok := v.(string)

	if !ok {
		return errors.New("could not cast value")
	}

	if err := bb.WriteUint(uint64(len([]rune(data))), bitbuffer.LittleEndian, bitsize); err != nil {
		return err
	}

	for i := 0; i < len(data); i++ {
		if err := bb.WriteByte(data[i]); err != nil {
			return err
		}
	}

	return nil
}

func marshalTimeOfDay(bb *bitbuffer.BitBuffer, v interface{}) error {
	tod, ok := v.(TimeOfDay)

	if !ok {
		return errors.New("could not cast value")
	}

	return bytecodec.MarshalToBitBuffer(bb, &tod)
}

func marshalDate(bb *bitbuffer.BitBuffer, v interface{}) error {
	date, ok := v.(Date)

	if !ok {
		return errors.New("could not cast value")
	}

	return bytecodec.MarshalToBitBuffer(bb, &date)
}

func marshalUTCTime(bb *bitbuffer.BitBuffer, v interface{}) error {
	utcTime, ok := v.(UTCTime)

	if !ok {
		return errors.New("could not cast value")
	}

	return bb.WriteUint(uint64(utcTime), bitbuffer.LittleEndian, 32)
}

func marshalClusterID(bb *bitbuffer.BitBuffer, v interface{}) error {
	clusterID, ok := v.(zigbee.ClusterID)

	if !ok {
		return errors.New("could not cast value")
	}

	return bb.WriteUint(uint64(clusterID), bitbuffer.LittleEndian, 16)
}

func marshalAttributeID(bb *bitbuffer.BitBuffer, v interface{}) error {
	attributeID, ok := v.(AttributeID)

	if !ok {
		return errors.New("could not cast value")
	}

	return bb.WriteUint(uint64(attributeID), bitbuffer.LittleEndian, 16)
}

func marshalIEEEAddress(bb *bitbuffer.BitBuffer, v interface{}) error {
	ieeeAddress, ok := v.(zigbee.IEEEAddress)

	if !ok {
		return errors.New("could not cast value")
	}

	return bb.WriteUint(uint64(ieeeAddress), bitbuffer.LittleEndian, 64)
}

func marshalSecurityKey(bb *bitbuffer.BitBuffer, v interface{}) error {
	networkKey, ok := v.(zigbee.NetworkKey)

	if !ok {
		return errors.New("could not cast value")
	}

	return bytecodec.MarshalToBitBuffer(bb, &networkKey)
}

func marshalBACnetOID(bb *bitbuffer.BitBuffer, v interface{}) error {
	oid, ok := v.(BACnetOID)

	if !ok {
		return errors.New("could not cast value")
	}

	return bb.WriteUint(uint64(oid), bitbuffer.LittleEndian, 32)
}

func marshalStructure(bb *bitbuffer.BitBuffer, v interface{}) error {
	values, ok := v.([]AttributeDataTypeValue)

	if !ok {
		return errors.New("could not cast value")
	}

	if err := bb.WriteUint(uint64(len(values)), bitbuffer.LittleEndian, 16); err != nil {
		return err
	}

	for _, val := range values {
		if err := val.Marshal(bb); err != nil {
			return err
		}
	}

	return nil
}

func marshalSlice(bb *bitbuffer.BitBuffer, v interface{}) error {
	slice, ok := v.(AttributeSlice)

	if !ok {
		return errors.New("could not cast value")
	}

	if err := bb.WriteByte(byte(slice.DataType)); err != nil {
		return err
	}

	if err := bb.WriteUint(uint64(len(slice.Values)), bitbuffer.LittleEndian, 16); err != nil {
		return err
	}

	for i := 0; i < len(slice.Values); i++ {
		if err := marshalZCLType(bb, slice.DataType, slice.Values[i]); err != nil {
			return err
		}
	}

	return nil
}

func marshalFloatSingle(bb *bitbuffer.BitBuffer, v interface{}) error {
	value, ok := v.(float32)

	if !ok {
		return errors.New("could not cast value")
	}

	bits := math.Float32bits(value)

	return bb.WriteUint(uint64(bits), bitbuffer.LittleEndian, 32)
}

func marshalFloatDouble(bb *bitbuffer.BitBuffer, v interface{}) error {
	value, ok := v.(float64)

	if !ok {
		return errors.New("could not cast value")
	}

	bits := math.Float64bits(value)

	return bb.WriteUint(bits, bitbuffer.LittleEndian, 64)
}

func (a *AttributeDataTypeValue) Unmarshal(bb *bitbuffer.BitBuffer) error {
	if dt, err := bb.ReadByte(); err != nil {
		return err
	} else {
		a.DataType = AttributeDataType(dt)
	}

	val, err := unmarshalZCLType(bb, a.DataType)

	if err != nil {
		return err
	}

	a.Value = val

	return nil
}

func unmarshalZCLType(bb *bitbuffer.BitBuffer, dt AttributeDataType) (interface{}, error) {
	switch dt {
	case TypeNull:
		return nil, nil
	case TypeData8:
		return unmarshalData(bb, 1)
	case TypeData16:
		return unmarshalData(bb, 2)
	case TypeData24:
		return unmarshalData(bb, 3)
	case TypeData32:
		return unmarshalData(bb, 4)
	case TypeData40:
		return unmarshalData(bb, 5)
	case TypeData48:
		return unmarshalData(bb, 6)
	case TypeData56:
		return unmarshalData(bb, 7)
	case TypeData64:
		return unmarshalData(bb, 8)
	case TypeBoolean:
		return unmarshalBoolean(bb)
	case TypeBitmap8:
		return unmarshalUint(bb, 8)
	case TypeBitmap16:
		return unmarshalUint(bb, 16)
	case TypeBitmap24:
		return unmarshalUint(bb, 24)
	case TypeBitmap32:
		return unmarshalUint(bb, 32)
	case TypeBitmap40:
		return unmarshalUint(bb, 40)
	case TypeBitmap48:
		return unmarshalUint(bb, 48)
	case TypeBitmap56:
		return unmarshalUint(bb, 56)
	case TypeBitmap64:
		return unmarshalUint(bb, 64)
	case TypeUnsignedInt8:
		return unmarshalUint(bb, 8)
	case TypeUnsignedInt16:
		return unmarshalUint(bb, 16)
	case TypeUnsignedInt24:
		return unmarshalUint(bb, 24)
	case TypeUnsignedInt32:
		return unmarshalUint(bb, 32)
	case TypeUnsignedInt40:
		return unmarshalUint(bb, 40)
	case TypeUnsignedInt48:
		return unmarshalUint(bb, 48)
	case TypeUnsignedInt56:
		return unmarshalUint(bb, 56)
	case TypeUnsignedInt64:
		return unmarshalUint(bb, 64)
	case TypeSignedInt8:
		return unmarshalInt(bb, 8)
	case TypeSignedInt16:
		return unmarshalInt(bb, 16)
	case TypeSignedInt24:
		return unmarshalInt(bb, 24)
	case TypeSignedInt32:
		return unmarshalInt(bb, 32)
	case TypeSignedInt40:
		return unmarshalInt(bb, 40)
	case TypeSignedInt48:
		return unmarshalInt(bb, 48)
	case TypeSignedInt56:
		return unmarshalInt(bb, 56)
	case TypeSignedInt64:
		return unmarshalInt(bb, 64)
	case TypeEnum8:
		val, err := unmarshalUint(bb, 8)

		if err == nil {
			val = uint8(val.(uint64))
		}

		return val, err
	case TypeEnum16:
		val, err := unmarshalUint(bb, 16)

		if err == nil {
			val = uint16(val.(uint64))
		}

		return val, err
	case TypeStringOctet8:
		return unmarshalString(bb, 8)
	case TypeStringOctet16:
		return unmarshalString(bb, 16)
	case TypeStringCharacter8:
		return unmarshalStringRune(bb, 8)
	case TypeStringCharacter16:
		return unmarshalStringRune(bb, 16)
	case TypeTimeOfDay:
		return unmarshalTimeOfDay(bb)
	case TypeDate:
		return unmarshalDate(bb)
	case TypeUTCTime:
		return unmarshalUTCTime(bb)
	case TypeClusterID:
		return unmarshalClusterID(bb)
	case TypeAttributeID:
		return unmarshalAttributeID(bb)
	case TypeIEEEAddress:
		return unmarshalIEEEAddress(bb)
	case TypeSecurityKey128:
		return unmarshalSecurityKey(bb)
	case TypeBACnetOID:
		return unmarshalBACnetOID(bb)
	case TypeStructure:
		return unmarshalStructure(bb)
	case TypeArray, TypeSet, TypeBag:
		return unmarshalSlice(bb)
	case TypeFloatSingle:
		return unmarshalFloatSingle(bb)
	case TypeFloatDouble:
		return unmarshalFloatDouble(bb)
	default:
		return nil, fmt.Errorf("unsupported ZCL type to unmarshal: %d", dt)
	}
}

func unmarshalData(bb *bitbuffer.BitBuffer, size int) (interface{}, error) {
	data := make([]byte, size)

	for i := size - 1; i >= 0; i-- {
		if b, err := bb.ReadByte(); err != nil {
			return nil, err
		} else {
			data[i] = b
		}
	}

	return data, nil
}

func unmarshalBoolean(bb *bitbuffer.BitBuffer) (interface{}, error) {
	if data, err := bb.ReadByte(); err != nil {
		return nil, err
	} else {
		return data != 0x00, nil
	}
}

func unmarshalUint(bb *bitbuffer.BitBuffer, bitsize int) (interface{}, error) {
	v, err := bb.ReadUint(bitbuffer.LittleEndian, bitsize)

	if err != nil {
		return nil, err
	}

	return v, nil
}

func unmarshalInt(bb *bitbuffer.BitBuffer, bitsize int) (interface{}, error) {
	v, err := bb.ReadInt(bitbuffer.LittleEndian, bitsize)

	if err != nil {
		return nil, err
	}

	return v, nil
}

func unmarshalString(bb *bitbuffer.BitBuffer, bitsize int) (interface{}, error) {
	if data, err := bb.ReadStringLengthPrefixed(bitbuffer.LittleEndian, bitsize); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func unmarshalStringRune(bb *bitbuffer.BitBuffer, bitsize int) (interface{}, error) {
	runeCount, err := bb.ReadUint(bitbuffer.LittleEndian, bitsize)

	if err != nil {
		return nil, err
	}

	var data []byte

	for i := 0; i < int(runeCount); i++ {
		b, err := bb.ReadByte()

		if err != nil {
			return nil, err
		}

		data = append(data, b)

		if b > 0x80 {
			moreBytes := 0

			if b < 0b1110000 {
				moreBytes = 1
			} else if b >= 0b1110000 && b < 0b11110000 {
				moreBytes = 2
			} else {
				moreBytes = 3
			}

			for i := 0; i < moreBytes; i++ {
				if b, err := bb.ReadByte(); err != nil {
					return nil, err
				} else {
					data = append(data, b)
				}
			}
		}
	}

	return string(data), nil
}

func unmarshalTimeOfDay(bb *bitbuffer.BitBuffer) (interface{}, error) {
	tod := TimeOfDay{}

	if err := bytecodec.UnmarshalFromBitBuffer(bb, &tod); err != nil {
		return nil, err
	}

	return tod, nil
}

func unmarshalDate(bb *bitbuffer.BitBuffer) (interface{}, error) {
	date := Date{}

	if err := bytecodec.UnmarshalFromBitBuffer(bb, &date); err != nil {
		return nil, err
	}

	return date, nil
}

func unmarshalUTCTime(bb *bitbuffer.BitBuffer) (interface{}, error) {
	v, err := bb.ReadInt(bitbuffer.LittleEndian, 32)

	if err != nil {
		return nil, err
	}

	return UTCTime(v), nil
}

func unmarshalClusterID(bb *bitbuffer.BitBuffer) (interface{}, error) {
	v, err := bb.ReadInt(bitbuffer.LittleEndian, 16)

	if err != nil {
		return nil, err
	}

	return zigbee.ClusterID(v), nil
}

func unmarshalAttributeID(bb *bitbuffer.BitBuffer) (interface{}, error) {
	v, err := bb.ReadInt(bitbuffer.LittleEndian, 16)

	if err != nil {
		return nil, err
	}

	return AttributeID(v), nil
}

func unmarshalIEEEAddress(bb *bitbuffer.BitBuffer) (interface{}, error) {
	v, err := bb.ReadInt(bitbuffer.LittleEndian, 64)

	if err != nil {
		return nil, err
	}

	return zigbee.IEEEAddress(v), nil
}

func unmarshalSecurityKey(bb *bitbuffer.BitBuffer) (interface{}, error) {
	networkKey := zigbee.NetworkKey{}

	if err := bytecodec.UnmarshalFromBitBuffer(bb, &networkKey); err != nil {
		return nil, err
	}

	return networkKey, nil
}

func unmarshalBACnetOID(bb *bitbuffer.BitBuffer) (interface{}, error) {
	v, err := bb.ReadInt(bitbuffer.LittleEndian, 32)

	if err != nil {
		return nil, err
	}

	return BACnetOID(v), nil
}

func unmarshalStructure(bb *bitbuffer.BitBuffer) (interface{}, error) {
	itemCount, err := bb.ReadUint(bitbuffer.LittleEndian, 16)

	if err != nil {
		return nil, err
	}

	values := []AttributeDataTypeValue{}

	for i := 0; i < int(itemCount); i++ {
		val := AttributeDataTypeValue{}

		if err := val.Unmarshal(bb); err != nil {
			return nil, err
		}

		values = append(values, val)
	}

	return values, nil
}

func unmarshalSlice(bb *bitbuffer.BitBuffer) (interface{}, error) {
	rawType, err := bb.ReadUint(bitbuffer.LittleEndian, 8)

	if err != nil {
		return nil, err
	}

	itemCount, err := bb.ReadUint(bitbuffer.LittleEndian, 16)

	if err != nil {
		return nil, err
	}

	itemType := AttributeDataType(rawType)

	value := AttributeSlice{
		DataType: itemType,
		Values:   []interface{}{},
	}

	for i := 0; i < int(itemCount); i++ {
		if val, err := unmarshalZCLType(bb, itemType); err != nil {
			return nil, err
		} else {
			value.Values = append(value.Values, val)
		}
	}

	return value, nil
}

func unmarshalFloatSingle(bb *bitbuffer.BitBuffer) (interface{}, error) {
	if bits, err := bb.ReadUint(bitbuffer.LittleEndian, 32); err != nil {
		return nil, err
	} else {
		return math.Float32frombits(uint32(bits)), nil
	}
}

func unmarshalFloatDouble(bb *bitbuffer.BitBuffer) (interface{}, error) {
	if bits, err := bb.ReadUint(bitbuffer.LittleEndian, 64); err != nil {
		return nil, err
	} else {
		return math.Float64frombits(bits), nil
	}
}

type BACnetOID uint32

type TimeOfDay struct {
	Hours      uint8
	Minutes    uint8
	Seconds    uint8
	Hundredths uint8
}

type Date struct {
	Year       uint8
	Month      uint8
	DayOfMonth uint8
	DayOfWeek  uint8
}

type UTCTime uint32
