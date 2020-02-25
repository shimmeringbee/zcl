package zcl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Unmarshal(t *testing.T) {
	t.Run("message with unknown command identifier", func(t *testing.T) {
		data := []byte{0b00000100, 0x20, 0x10, 0x40, 0xff}

		_, err := Unmarshal(data)
		assert.Error(t, err)
	})

	t.Run("ReadAttributes with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   ReadAttributesID,
			},
			Command: &ReadAttributes{
				Identifier: []AttributeIdentifier{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("ReadAttributesResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   ReadAttributesResponseID,
			},
			Command: &ReadAttributesResponse{
				Records: []ReadAttributeResponseRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("WriteAttributes with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   WriteAttributesID,
			},
			Command: &WriteAttributes{
				Records: []WriteAttributesRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("WriteAttributesResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   WriteAttributesResponseID,
			},
			Command: &WriteAttributesResponse{
				Records: []WriteAttributesResponseRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("WriteAttributesNoResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   WriteAttributesNoResponseID,
			},
			Command: &WriteAttributesNoResponse{
				Records: []WriteAttributesRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("WriteAttributesUndivided with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   WriteAttributesUndividedID,
			},
			Command: &WriteAttributesUndivided{
				Records: []WriteAttributesRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("ReportAttributes with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   ReportAttributesID,
			},
			Command: &ReportAttributes{
				Records: []ReportAttributesRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DefaultResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DefaultResponseID,
			},
			Command: &DefaultResponse{
				CommandIdentifier: 0xaa,
				Status:            0x01,
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DiscoverAttributes with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DiscoverAttributesID,
			},
			Command: &DiscoverAttributes{},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DiscoverAttributesResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DiscoverAttributesResponseID,
			},
			Command: &DiscoverAttributesResponse{
				DiscoveryComplete: false,
				Records:           []DiscoverAttributesResponseRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("ReadAttributesStructured with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   ReadAttributesStructuredID,
			},
			Command: &ReadAttributesStructured{
				Records: []ReadAttributesStructuredRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("ReadAttributesStructured with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   ReadAttributesStructuredID,
			},
			Command: &ReadAttributesStructured{
				Records: []ReadAttributesStructuredRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("WriteAttributesStructured with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   WriteAttributesStructuredID,
			},
			Command: &WriteAttributesStructured{
				Records: []WriteAttributesStructuredRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("WriteAttributesStructuredResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   WriteAttributesStructuredResponseID,
			},
			Command: &WriteAttributesStructuredResponse{
				Records: []WriteAttributesStructuredResponseRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DiscoverCommandsReceived with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DiscoverCommandsReceivedID,
			},
			Command: &DiscoverCommandsReceived{},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DiscoverCommandsReceivedResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DiscoverCommandsReceivedResponseID,
			},
			Command: &DiscoverCommandsReceivedResponse{
				CommandIdentifier: []uint8{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DiscoverCommandsGenerated with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DiscoverCommandsGeneratedID,
			},
			Command: &DiscoverCommandsGenerated{},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DiscoverCommandsGeneratedResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DiscoverCommandsGeneratedResponseID,
			},
			Command: &DiscoverCommandsGeneratedResponse{
				CommandIdentifier: []uint8{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DiscoverAttributesExtended with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DiscoverAttributesExtendedID,
			},
			Command: &DiscoverAttributesExtended{},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})

	t.Run("DiscoverAttributesExtendedResponse with manufacturer specific", func(t *testing.T) {
		expectedMessage := ZCLFrame{
			Header: Header{
				Control: Control{
					Reserved:               0,
					DisableDefaultResponse: false,
					Direction:              ClientToServer,
					ManufacturerSpecific:   true,
					FrameType:              FrameGlobal,
				},
				Manufacturer:        0x1020,
				TransactionSequence: 0x40,
				CommandIdentifier:   DiscoverAttributesExtendedResponseID,
			},
			Command: &DiscoverAttributesExtendedResponse{
				Records: []DiscoverAttributesExtendedResponseRecord{},
			},
		}

		bytes, err := Marshal(expectedMessage)
		assert.NoError(t, err)

		actualMessage, err := Unmarshal(bytes)

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, actualMessage)
	})
}
