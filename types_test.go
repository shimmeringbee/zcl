package zcl

import (
	"github.com/shimmeringbee/bytecodec"
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
}
