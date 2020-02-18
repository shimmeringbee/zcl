package zcl

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadAttributes(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadAttributes{
			Identifier: []AttributeIdentifier{
				0x1020,
			},
		}
		actualCommand := ReadAttributes{}
		expectedBytes := []byte{0x20, 0x10}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_ReadAttributesResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadAttributesResponse{
			Records: []ReadAttributeResponseStatusRecord{
				{
					Identifier: 0x1000,
					Status:     0,
					DataTypeValue: &AttributeDataTypeValue{
						DataType: TypeData8,
						Value:    []byte{0xaa},
					},
				},
				{
					Identifier:    0x2000,
					Status:        0x01,
					DataTypeValue: nil,
				},
			},
		}
		actualCommand := ReadAttributesResponse{}
		expectedBytes := []byte{0x00, 0x10, 0x00, 0x08, 0xaa, 0x00, 0x20, 0x01}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}
