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
}
