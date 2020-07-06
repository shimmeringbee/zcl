package zcl

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AttributeDataTypeValue(t *testing.T) {
	t.Run("marshaling and unmarshaling of unsupported type", func(t *testing.T) {
		inputValue := &AttributeDataTypeValue{
			DataType: TypeUnknown,
			Value:    nil,
		}

		inputBytes := []byte{0xff}

		_, err := bytecodec.Marshal(&inputValue)
		assert.Error(t, err)

		err = bytecodec.Unmarshal(inputBytes, &AttributeDataTypeValue{})
		assert.Error(t, err)
	})

	t.Run("marshaling and unmarshaling of null type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeNull,
			Value:    nil,
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x00}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of data8 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeData8,
			Value:    []byte{0xaa},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x08, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of data16 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeData16,
			Value:    []byte{0x11, 0x22},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x09, 0x22, 0x11}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of data24 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeData24,
			Value:    []byte{0x11, 0x22, 0x33},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x0a, 0x33, 0x22, 0x11}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of data32 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeData32,
			Value:    []byte{0x11, 0x22, 0x33, 0x44},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x0b, 0x44, 0x33, 0x22, 0x11}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of data40 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeData40,
			Value:    []byte{0x11, 0x22, 0x33, 0x44, 0x55},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x0c, 0x55, 0x44, 0x33, 0x22, 0x11}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of data48 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeData48,
			Value:    []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x0d, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of data56 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeData56,
			Value:    []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x0e, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of data64 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeData64,
			Value:    []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x0f, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of boolean type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBoolean,
			Value:    true,
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x10, 0x01}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of uint8 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeUnsignedInt8,
			Value:    uint64(0xaa),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x20, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of uint16 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeUnsignedInt16,
			Value:    uint64(0xaadd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x21, 0xdd, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of uint24 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeUnsignedInt24,
			Value:    uint64(0xaa00dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x22, 0xdd, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of uint32 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeUnsignedInt32,
			Value:    uint64(0xaa0001dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x23, 0xdd, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of uint40 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeUnsignedInt40,
			Value:    uint64(0xaa000102dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x24, 0xdd, 0x02, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of uint48 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeUnsignedInt48,
			Value:    uint64(0xaa00010203dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x25, 0xdd, 0x03, 0x02, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of uint56 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeUnsignedInt56,
			Value:    uint64(0xaa0001020304dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x26, 0xdd, 0x04, 0x03, 0x02, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of uint64 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeUnsignedInt64,
			Value:    uint64(0xaa000102030405dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x27, 0xdd, 0x05, 0x04, 0x03, 0x02, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of bitmap8 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBitmap8,
			Value:    uint64(0xaa),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x18, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of bitmap16 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBitmap16,
			Value:    uint64(0xaadd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x19, 0xdd, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of bitmap24 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBitmap24,
			Value:    uint64(0xaa00dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x1a, 0xdd, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of bitmap32 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBitmap32,
			Value:    uint64(0xaa0001dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x1b, 0xdd, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of bitmap40 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBitmap40,
			Value:    uint64(0xaa000102dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x1c, 0xdd, 0x02, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of bitmap48 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBitmap48,
			Value:    uint64(0xaa00010203dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x1d, 0xdd, 0x03, 0x02, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of bitmap56 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBitmap56,
			Value:    uint64(0xaa0001020304dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x1e, 0xdd, 0x04, 0x03, 0x02, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of bitmap64 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBitmap64,
			Value:    uint64(0xaa000102030405dd),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x1f, 0xdd, 0x05, 0x04, 0x03, 0x02, 0x01, 0x00, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of int8 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSignedInt8,
			Value:    int64(-128),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x28, 0x80}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of int16 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSignedInt16,
			Value:    int64(-128),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x29, 0x80, 0xff}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of int24 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSignedInt24,
			Value:    int64(-128),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x2a, 0x80, 0xff, 0xff}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of int32 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSignedInt32,
			Value:    int64(-128),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x2b, 0x80, 0xff, 0xff, 0xff}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of int40 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSignedInt40,
			Value:    int64(-128),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x2c, 0x80, 0xff, 0xff, 0xff, 0xff}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of int48 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSignedInt48,
			Value:    int64(-128),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x2d, 0x80, 0xff, 0xff, 0xff, 0xff, 0xff}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of int56 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSignedInt56,
			Value:    int64(-128),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x2e, 0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of int64 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSignedInt64,
			Value:    int64(-128),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x2f, 0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of enum8 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeEnum8,
			Value:    uint8(0x01),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x30, 0x01}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of enum16 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeEnum16,
			Value:    uint16(0x0001),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x31, 0x01, 0x00}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of octet 8 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeStringOctet8,
			Value:    "Hi",
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x41, 0x02, 'H', 'i'}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of octet 16 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeStringOctet16,
			Value:    "Hi",
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x43, 0x02, 0x00, 'H', 'i'}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of time of day", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeTimeOfDay,
			Value: TimeOfDay{
				Hours:      1,
				Minutes:    2,
				Seconds:    3,
				Hundredths: 4,
			},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0xe0, 0x01, 0x02, 0x03, 0x04}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of date", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeDate,
			Value: Date{
				Year:       1,
				Month:      2,
				DayOfMonth: 3,
				DayOfWeek:  4,
			},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0xe1, 0x01, 0x02, 0x03, 0x04}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of utc time", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeUTCTime,
			Value:    UTCTime(0x01020304),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0xe2, 0x04, 0x03, 0x02, 0x01}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of cluster ID", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeClusterID,
			Value:    zigbee.ClusterID(0x0102),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0xe9, 0x02, 0x01}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of attribute ID", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeAttributeID,
			Value:    AttributeID(0x0102),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0xea, 0x02, 0x01}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of IEEE address", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeIEEEAddress,
			Value:    zigbee.IEEEAddress(0x0102030405060708),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0xf0, 0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of IEEE address", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSecurityKey128,
			Value:    zigbee.NetworkKey{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0xf1, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of BACnet OID", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBACnetOID,
			Value:    BACnetOID(0x01020304),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0xeb, 0x04, 0x03, 0x02, 0x01}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of structure", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeStructure,
			Value: []AttributeDataTypeValue{
				{
					DataType: TypeUnsignedInt8,
					Value:    uint64(0x01),
				},
				{
					DataType: TypeUnsignedInt8,
					Value:    uint64(0x02),
				},
			},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x4c, 0x02, 0x00, 0x20, 0x01, 0x20, 0x02}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of array", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeArray,
			Value: AttributeSlice{
				DataType: TypeUnsignedInt8,
				Values:   []interface{}{uint64(0x01), uint64(0x02), uint64(0x03)},
			},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x48, 0x20, 0x03, 0x00, 0x01, 0x02, 0x03}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of set", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeSet,
			Value: AttributeSlice{
				DataType: TypeUnsignedInt8,
				Values:   []interface{}{uint64(0x01), uint64(0x02), uint64(0x03)},
			},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x50, 0x20, 0x03, 0x00, 0x01, 0x02, 0x03}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of bag", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeBag,
			Value: AttributeSlice{
				DataType: TypeUnsignedInt8,
				Values:   []interface{}{uint64(0x01), uint64(0x02), uint64(0x03)},
			},
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x51, 0x20, 0x03, 0x00, 0x01, 0x02, 0x03}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of single precision floating point", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeFloatSingle,
			Value:    float32(9.807),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x39, 0x79, 0xe9, 0x1c, 0x41}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of double precision floating point", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeFloatDouble,
			Value:    float64(9.807),
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x3a, 0x77, 0xbe, 0x9f, 0x1a, 0x2f, 0x9d, 0x23, 0x40}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of character 8 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeStringCharacter8,
			Value:    "Zig🐝",
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x42, 0x04, 0x5a, 0x69, 0x67, 0xf0, 0x9f, 0x90, 0x9d}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of character 16 type", func(t *testing.T) {
		expectedValue := &AttributeDataTypeValue{
			DataType: TypeStringCharacter16,
			Value:    "Zig🐝",
		}
		actualValue := &AttributeDataTypeValue{}
		expectedBytes := []byte{0x44, 0x04, 0x00, 0x5a, 0x69, 0x67, 0xf0, 0x9f, 0x90, 0x9d}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})
}

func Test_AttributeDataValue(t *testing.T) {
	t.Run("marshaling and unmarshaling of attribute data value with prior type", func(t *testing.T) {
		type SUT struct {
			DataType AttributeDataType
			One      *AttributeDataValue
			Two      *AttributeDataValue
		}

		expectedValue := SUT{
			DataType: TypeStringCharacter8,
			One:      &AttributeDataValue{Value: "One"},
			Two:      &AttributeDataValue{Value: "Two"},
		}

		actualValue := SUT{}
		expectedBytes := []byte{0x42, 0x03, 'O', 'n', 'e', 0x03, 'T', 'w', 'o'}

		actualBytes, err := bytecodec.Marshal(&expectedValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("marshaling and unmarshaling of attribute data value without prior type errors", func(t *testing.T) {
		type SUT struct {
			One      *AttributeDataValue
			Two      *AttributeDataValue
			DataType AttributeDataType
		}

		expectedValue := SUT{
			One:      &AttributeDataValue{Value: "One"},
			Two:      &AttributeDataValue{Value: "Two"},
			DataType: TypeStringCharacter8,
		}

		_, err := bytecodec.Marshal(&expectedValue)
		assert.Error(t, err)
	})
}
