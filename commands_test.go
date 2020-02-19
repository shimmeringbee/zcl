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
			Records: []ReadAttributeResponseRecord{
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

func Test_WriteAttributes(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := WriteAttributes{
			Records: []WriteAttributesRecord{
				{
					Identifier: 0x1020,
					DataTypeValue: &AttributeDataTypeValue{
						DataType: TypeData8,
						Value:    []byte{0xaa},
					},
				},
			},
		}
		actualCommand := WriteAttributes{}
		expectedBytes := []byte{0x20, 0x10, 0x08, 0x0aa}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_WriteAttributesResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := WriteAttributesResponse{
			Records: []WriteAttributesResponseRecord{
				{
					Status:     0x01,
					Identifier: 0x1020,
				},
			},
		}
		actualCommand := WriteAttributesResponse{}
		expectedBytes := []byte{0x01, 0x20, 0x10}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_WriteAttributesNoResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := WriteAttributesNoResponse{
			Records: []WriteAttributesRecord{
				{
					Identifier: 0x1020,
					DataTypeValue: &AttributeDataTypeValue{
						DataType: TypeData8,
						Value:    []byte{0xaa},
					},
				},
			},
		}
		actualCommand := WriteAttributesNoResponse{}
		expectedBytes := []byte{0x20, 0x10, 0x08, 0x0aa}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}
