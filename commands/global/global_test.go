package global

import (
	"github.com/shimmeringbee/bytecodec"
	"github.com/shimmeringbee/zcl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadAttributes(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadAttributes{
			Identifier: []zcl.AttributeID{
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&ReadAttributes{})
		assert.NoError(t, err)
		assert.Equal(t, ReadAttributesID, id)
	})
}

func Test_ReadAttributesResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadAttributesResponse{
			Records: []ReadAttributeResponseRecord{
				{
					Identifier: 0x1000,
					Status:     0,
					DataTypeValue: &zcl.AttributeDataTypeValue{
						DataType: zcl.TypeData8,
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&ReadAttributesResponse{})
		assert.NoError(t, err)
		assert.Equal(t, ReadAttributesResponseID, id)
	})
}

func Test_WriteAttributes(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := WriteAttributes{
			Records: []WriteAttributesRecord{
				{
					Identifier: 0x1020,
					DataTypeValue: &zcl.AttributeDataTypeValue{
						DataType: zcl.TypeData8,
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&WriteAttributes{})
		assert.NoError(t, err)
		assert.Equal(t, WriteAttributesID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&WriteAttributesResponse{})
		assert.NoError(t, err)
		assert.Equal(t, WriteAttributesResponseID, id)
	})
}

func Test_WriteAttributesNoResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := WriteAttributesNoResponse{
			Records: []WriteAttributesRecord{
				{
					Identifier: 0x1020,
					DataTypeValue: &zcl.AttributeDataTypeValue{
						DataType: zcl.TypeData8,
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&WriteAttributesNoResponse{})
		assert.NoError(t, err)
		assert.Equal(t, WriteAttributesNoResponseID, id)
	})
}

func Test_WriteAttributesUndivided(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := WriteAttributesUndivided{
			Records: []WriteAttributesRecord{
				{
					Identifier: 0x1020,
					DataTypeValue: &zcl.AttributeDataTypeValue{
						DataType: zcl.TypeData8,
						Value:    []byte{0xaa},
					},
				},
			},
		}
		actualCommand := WriteAttributesUndivided{}
		expectedBytes := []byte{0x20, 0x10, 0x08, 0x0aa}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&WriteAttributesUndivided{})
		assert.NoError(t, err)
		assert.Equal(t, WriteAttributesUndividedID, id)
	})
}

func Test_ConfigureReporting(t *testing.T) {
	t.Run("direction 0x00, with discrete type, marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ConfigureReporting{
			Records: []ConfigureReportingRecord{
				{
					Direction:        0x00,
					Identifier:       0x1020,
					DataType:         zcl.TypeData8,
					MinimumInterval:  0x3040,
					MaximumInterval:  0x5060,
					ReportableChange: &zcl.AttributeDataValue{},
					Timeout:          0,
				},
			},
		}
		actualCommand := ConfigureReporting{}
		expectedBytes := []byte{0x00, 0x20, 0x10, 0x08, 0x40, 0x30, 0x60, 0x50}

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
					DataType:         zcl.TypeUnsignedInt8,
					MinimumInterval:  0x3040,
					MaximumInterval:  0x5060,
					ReportableChange: &zcl.AttributeDataValue{Value: uint64(0xaa)},
					Timeout:          0,
				},
			},
		}
		actualCommand := ConfigureReporting{}
		expectedBytes := []byte{0x00, 0x20, 0x10, 0x20, 0x40, 0x30, 0x60, 0x50, 0xaa}

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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&ConfigureReporting{})
		assert.NoError(t, err)
		assert.Equal(t, ConfigureReportingID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&ConfigureReportingResponse{})
		assert.NoError(t, err)
		assert.Equal(t, ConfigureReportingResponseID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&ReadReportingConfiguration{})
		assert.NoError(t, err)
		assert.Equal(t, ReadReportingConfigurationID, id)
	})
}

func Test_ReadReportingConfigurationResponse(t *testing.T) {
	t.Run("direction 0x00, with discrete type, marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReadReportingConfigurationResponse{
			Records: []ReadReportingConfigurationResponseRecord{
				{
					Status:           0x99,
					Direction:        0x00,
					Identifier:       0x1020,
					DataType:         zcl.TypeData8,
					MinimumInterval:  0x3040,
					MaximumInterval:  0x5060,
					ReportableChange: &zcl.AttributeDataValue{},
					Timeout:          0,
				},
			},
		}
		actualCommand := ReadReportingConfigurationResponse{}
		expectedBytes := []byte{0x99, 0x00, 0x20, 0x10, 0x08, 0x40, 0x30, 0x60, 0x50}

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
					DataType:         zcl.TypeUnsignedInt8,
					MinimumInterval:  0x3040,
					MaximumInterval:  0x5060,
					ReportableChange: &zcl.AttributeDataValue{Value: uint64(0xaa)},
					Timeout:          0,
				},
			},
		}
		actualCommand := ReadReportingConfigurationResponse{}
		expectedBytes := []byte{0x99, 0x00, 0x20, 0x10, 0x20, 0x40, 0x30, 0x60, 0x50, 0xaa}

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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&ReadReportingConfigurationResponse{})
		assert.NoError(t, err)
		assert.Equal(t, ReadReportingConfigurationResponseID, id)
	})
}

func Test_ReportAttributes(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := ReportAttributes{
			Records: []ReportAttributesRecord{
				{
					Identifier: 0x1020,
					DataTypeValue: &zcl.AttributeDataTypeValue{
						DataType: zcl.TypeData8,
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&ReportAttributes{})
		assert.NoError(t, err)
		assert.Equal(t, ReportAttributesID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&DefaultResponse{})
		assert.NoError(t, err)
		assert.Equal(t, DefaultResponseID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&DiscoverAttributes{})
		assert.NoError(t, err)
		assert.Equal(t, DiscoverAttributesID, id)
	})
}

func Test_DiscoverAttributesResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DiscoverAttributesResponse{
			DiscoveryComplete: true,
			Records: []DiscoverAttributesResponseRecord{
				{
					Identifier: 0x1020,
					DataType:   zcl.TypeData8,
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&DiscoverAttributesResponse{})
		assert.NoError(t, err)
		assert.Equal(t, DiscoverAttributesResponseID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&ReadAttributesStructured{})
		assert.NoError(t, err)
		assert.Equal(t, ReadAttributesStructuredID, id)
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
					DataTypeValue: &zcl.AttributeDataTypeValue{
						DataType: zcl.TypeData8,
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&WriteAttributesStructured{})
		assert.NoError(t, err)
		assert.Equal(t, WriteAttributesStructuredID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&WriteAttributesStructuredResponse{})
		assert.NoError(t, err)
		assert.Equal(t, WriteAttributesStructuredResponseID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&DiscoverCommandsReceived{})
		assert.NoError(t, err)
		assert.Equal(t, DiscoverCommandsReceivedID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&DiscoverCommandsReceivedResponse{})
		assert.NoError(t, err)
		assert.Equal(t, DiscoverCommandsReceivedResponseID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&DiscoverCommandsGenerated{})
		assert.NoError(t, err)
		assert.Equal(t, DiscoverCommandsGeneratedID, id)
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

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&DiscoverCommandsGeneratedResponse{})
		assert.NoError(t, err)
		assert.Equal(t, DiscoverCommandsGeneratedResponseID, id)
	})
}

func Test_DiscoverAttributesExtended(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DiscoverAttributesExtended{
			StartAttributeIdentifier:  0x1020,
			MaximumNumberOfAttributes: 0x30,
		}
		actualCommand := DiscoverAttributesExtended{}
		expectedBytes := []byte{0x20, 0x10, 0x30}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&DiscoverAttributesExtended{})
		assert.NoError(t, err)
		assert.Equal(t, DiscoverAttributesExtendedID, id)
	})
}

func Test_DiscoverAttributesExtendedResponse(t *testing.T) {
	t.Run("marshals and unmarshals correctly", func(t *testing.T) {
		expectedCommand := DiscoverAttributesExtendedResponse{
			DiscoveryComplete: true,
			Records: []DiscoverAttributesExtendedResponseRecord{
				{
					Identifier:    0x1020,
					DataType:      zcl.TypeData8,
					AccessControl: 0x04,
				},
			},
		}
		actualCommand := DiscoverAttributesExtendedResponse{}
		expectedBytes := []byte{0x01, 0x20, 0x10, 0x08, 0x04}

		actualBytes, err := bytecodec.Marshal(&expectedCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, actualBytes)

		err = bytecodec.Unmarshal(expectedBytes, &actualCommand)
		assert.NoError(t, err)
		assert.Equal(t, expectedCommand, actualCommand)
	})

	t.Run("the message is registered in the command registry", func(t *testing.T) {
		cr := zcl.NewCommandRegistry()
		Register(cr)

		id, err := cr.GetGlobalCommandIdentifier(&DiscoverAttributesExtendedResponse{})
		assert.NoError(t, err)
		assert.Equal(t, DiscoverAttributesExtendedResponseID, id)
	})
}
