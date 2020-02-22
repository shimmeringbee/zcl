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

func Test_ConfigureReporting(t *testing.T) {
	t.Skip("skipping ConfigureReporting due to marshaler/unmarshaler being unable to handle order of attributes")

	t.Run("direction 0x00, with discrete type, marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ConfigureReporting{
			Records: []ConfigureReportingRecord{
				{
					Direction:        0x00,
					Identifier:       0x1020,
					DataType:         TypeData8,
					MinimumInterval:  0x3040,
					MaximumInterval:  0x5060,
					ReportableChange: nil,
					Timeout:          0x7080,
				},
			},
		}
		actualCommand := ConfigureReporting{}
		expectedBytes := []byte{0x00, 0x20, 0x10, 0x08, 0x40, 0x30, 0x60, 0x50, 0x80, 0x70}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("direction 0x00, with analogue type, marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ConfigureReporting{
			Records: []ConfigureReportingRecord{
				{
					Direction:        0x00,
					Identifier:       0x1020,
					DataType:         TypeUnsignedInt8,
					MinimumInterval:  0x3040,
					MaximumInterval:  0x5060,
					ReportableChange: 0xaa,
					Timeout:          0x7080,
				},
			},
		}
		actualCommand := ConfigureReporting{}
		expectedBytes := []byte{0x00, 0x20, 0x10, 0x08, 0x40, 0x30, 0x60, 0x50, 0xaa, 0x80, 0x70}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("direction 0x01, marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ConfigureReporting{
			Records: []ConfigureReportingRecord{
				{
					Direction:        0x01,
					Identifier:       0x1020,
					DataType:         0,
					MinimumInterval:  0,
					MaximumInterval:  0,
					ReportableChange: nil,
					Timeout:          0x3040,
				},
			},
		}
		actualCommand := ConfigureReporting{}
		expectedBytes := []byte{0x01, 0x20, 0x10, 0x40, 0x30}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_ConfigureReportingResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ConfigureReportingResponse{
			Records: []ConfigureReportingResponseRecord{
				{
					Status:     0x01,
					Direction:  0x02,
					Identifier: 0x1020,
				},
			},
		}
		actualCommand := ConfigureReportingResponse{}
		expectedBytes := []byte{0x01, 0x02, 0x20, 0x10}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_ReadReportingConfiguration(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadReportingConfiguration{
			Records: []ReadReportingConfigurationRecord{
				{
					Direction:  0x02,
					Identifier: 0x1020,
				},
			},
		}
		actualCommand := ReadReportingConfiguration{}
		expectedBytes := []byte{0x02, 0x20, 0x10}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_ReadReportingConfigurationResponse(t *testing.T) {
	t.Skip("skipping ReadReportingConfigurationResponse due to marshaler/unmarshaler being unable to handle order of attributes")

	t.Run("direction 0x00, with discrete type, marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadReportingConfigurationResponse{
			Records: []ReadReportingConfigurationResponseRecord{
				{
					Status:           0x99,
					Direction:        0x00,
					Identifier:       0x1020,
					DataType:         TypeData8,
					MinimumInterval:  0x3040,
					MaximumInterval:  0x5060,
					ReportableChange: nil,
					Timeout:          0x7080,
				},
			},
		}
		actualCommand := ReadReportingConfigurationResponse{}
		expectedBytes := []byte{0x99, 0x00, 0x20, 0x10, 0x08, 0x40, 0x30, 0x60, 0x50, 0x80, 0x70}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("direction 0x00, with analogue type, marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadReportingConfigurationResponse{
			Records: []ReadReportingConfigurationResponseRecord{
				{
					Status:           0x99,
					Direction:        0x00,
					Identifier:       0x1020,
					DataType:         TypeUnsignedInt8,
					MinimumInterval:  0x3040,
					MaximumInterval:  0x5060,
					ReportableChange: 0xaa,
					Timeout:          0x7080,
				},
			},
		}
		actualCommand := ReadReportingConfigurationResponse{}
		expectedBytes := []byte{0x99, 0x00, 0x20, 0x10, 0x08, 0x40, 0x30, 0x60, 0x50, 0xaa, 0x80, 0x70}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("direction 0x01, marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadReportingConfigurationResponse{
			Records: []ReadReportingConfigurationResponseRecord{
				{
					Status:           0x99,
					Direction:        0x01,
					Identifier:       0x1020,
					DataType:         0,
					MinimumInterval:  0,
					MaximumInterval:  0,
					ReportableChange: nil,
					Timeout:          0x3040,
				},
			},
		}
		actualCommand := ReadReportingConfigurationResponse{}
		expectedBytes := []byte{0x99, 0x01, 0x20, 0x10, 0x40, 0x30}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_ReportAttributes(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReportAttributes{
			Records: []ReportAttributesRecord{
				{
					Identifier: 0x1020,
					DataTypeValue: &AttributeDataTypeValue{
						DataType: TypeData8,
						Value:    []byte{0xaa},
					},
				},
			},
		}
		actualCommand := ReportAttributes{}
		expectedBytes := []byte{0x20, 0x10, 0x08, 0x0aa}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_DefaultResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DefaultResponse{
			CommandIdentifier: 0x01,
			Status:            0x02,
		}
		actualCommand := DefaultResponse{}
		expectedBytes := []byte{0x01, 0x02}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_DiscoverAttributes(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DiscoverAttributes{
			StartAttributeIdentifier:  0x1020,
			MaximumNumberOfAttributes: 0x40,
		}
		actualCommand := DiscoverAttributes{}
		expectedBytes := []byte{0x20, 0x10, 0x40}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_DiscoverAttributesResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DiscoverAttributesResponse{
			DiscoveryComplete: true,
			Records: []DiscoverAttributesResponseRecord{
				{
					Identifier: 0x1020,
					DataType:   TypeData8,
				},
			},
		}
		actualCommand := DiscoverAttributesResponse{}
		expectedBytes := []byte{0x01, 0x20, 0x10, 0x08}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_ReadAttributesStructured(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadAttributesStructured{
			Records: []ReadAttributesStructuredRecord{
				{
					Identifier: 0x1020,
					Selector: Selector{
						Index: []uint16{0x3040, 0x5060},
					},
				},
			},
		}
		actualCommand := ReadAttributesStructured{}
		expectedBytes := []byte{0x20, 0x10, 0x02, 0x40, 0x30, 0x60, 0x50}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_WriteAttributesStructured(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := WriteAttributesStructured{
			Records: []WriteAttributesStructuredRecord{
				{
					Identifier: 0x1020,
					Selector: Selector{
						BagSetOperation: 0x02,
						Index:           []uint16{0x3040, 0x5060},
					},
					DataTypeValue: &AttributeDataTypeValue{
						DataType: TypeData8,
						Value:    []byte{0xaa},
					},
				},
			},
		}
		actualCommand := WriteAttributesStructured{}
		expectedBytes := []byte{0x20, 0x10, 0x22, 0x40, 0x30, 0x60, 0x50, 0x08, 0xaa}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_WriteAttributesStructuredResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := WriteAttributesStructuredResponse{
			Records: []WriteAttributesStructuredResponseRecord{
				{
					Status:     0x01,
					Identifier: 0x1020,
					Selector: Selector{
						BagSetOperation: 0x02,
						Index:           []uint16{0x3040, 0x5060},
					},
				},
			},
		}
		actualCommand := WriteAttributesStructuredResponse{}
		expectedBytes := []byte{0x01, 0x20, 0x10, 0x22, 0x40, 0x30, 0x60, 0x50}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_DiscoverCommandsReceived(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DiscoverCommandsReceived{
			StartCommandIdentifier:  0x10,
			MaximumNumberOfCommands: 0x40,
		}
		actualCommand := DiscoverCommandsReceived{}
		expectedBytes := []byte{0x10, 0x40}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_DiscoverCommandsReceivedResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DiscoverCommandsReceivedResponse{
			DiscoveryComplete: true,
			CommandIdentifier: []uint8{0x02, 0x3},
		}
		actualCommand := DiscoverCommandsReceivedResponse{}
		expectedBytes := []byte{0x01, 0x02, 0x03}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_DiscoverCommandsGenerated(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DiscoverCommandsGenerated{
			StartCommandIdentifier:  0x10,
			MaximumNumberOfCommands: 0x40,
		}
		actualCommand := DiscoverCommandsGenerated{}
		expectedBytes := []byte{0x10, 0x40}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}

func Test_DiscoverCommandsGeneratedResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DiscoverCommandsGeneratedResponse{
			DiscoveryComplete: true,
			CommandIdentifier: []uint8{0x02, 0x3},
		}
		actualCommand := DiscoverCommandsGeneratedResponse{}
		expectedBytes := []byte{0x01, 0x02, 0x03}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})
}
