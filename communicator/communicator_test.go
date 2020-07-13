package communicator

import (
	"context"
	"fmt"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zcl/commands/global"
	"github.com/shimmeringbee/zigbee"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCommunicator_ProcessIncomingMessage(t *testing.T) {
	t.Run("an unrecognised/unparsable message results in an error", func(t *testing.T) {
		provider := &zigbee.MockProvider{}
		cr := zcl.NewCommandRegistry()

		c := NewCommunicator(provider, cr)

		err := c.ProcessIncomingMessage(zigbee.NodeIncomingMessageEvent{
			Node: zigbee.Node{},
			IncomingMessage: zigbee.IncomingMessage{
				GroupID:              0,
				SourceIEEEAddress:    0,
				SourceNetworkAddress: 0,
				Broadcast:            false,
				Secure:               false,
				LinkQuality:          0,
				Sequence:             0,
				ApplicationMessage: zigbee.ApplicationMessage{
					ClusterID:           0,
					SourceEndpoint:      0,
					DestinationEndpoint: 0,
					Data:                []byte{},
				},
			},
		})

		assert.Error(t, err)
		fmt.Println(err)
	})
}

func TestCommunicator_Request(t *testing.T) {
	t.Run("sending a message via the communicator goes to provider", func(t *testing.T) {
		provider := &zigbee.MockProvider{}
		cr := zcl.NewCommandRegistry()
		global.Register(cr)

		c := NewCommunicator(provider, cr)

		expectedIEEE := zigbee.IEEEAddress(2)
		message := zcl.Message{
			FrameType:           zcl.FrameGlobal,
			Direction:           zcl.ServerToClient,
			TransactionSequence: 255,
			Manufacturer:        0,
			ClusterID:           0x0001,
			SourceEndpoint:      1,
			DestinationEndpoint: 2,
			Command: &global.ReadAttributesResponse{
				Records: []global.ReadAttributeResponseRecord{
					{
						Identifier: 1,
						Status:     0,
						DataTypeValue: &zcl.AttributeDataTypeValue{
							DataType: zcl.TypeSignedInt8,
							Value:    int64(64),
						},
					},
				},
			},
		}

		expectedAppMessage, _ := cr.Marshal(message)

		provider.On("SendApplicationMessageToNode", mock.Anything, expectedIEEE, expectedAppMessage, false).Return(nil)

		err := c.Request(context.Background(), expectedIEEE, false, message)
		assert.NoError(t, err)

		provider.AssertExpectations(t)
	})
}

func TestCommunicator_RequestResponse(t *testing.T) {
	t.Run("a message via the communicator goes to provider and the reply is returned to the caller, as success", func(t *testing.T) {
		provider := &zigbee.MockProvider{}
		cr := zcl.NewCommandRegistry()
		global.Register(cr)

		c := NewCommunicator(provider, cr)

		expectedIEEE := zigbee.IEEEAddress(2)
		requestMessage := zcl.Message{
			FrameType:           zcl.FrameGlobal,
			Direction:           zcl.ClientToServer,
			TransactionSequence: 255,
			Manufacturer:        0,
			ClusterID:           0x0001,
			SourceEndpoint:      1,
			DestinationEndpoint: 2,
			Command: &global.ReadAttributes{
				Identifier: []zcl.AttributeID{1},
			},
		}

		expectedAppMessage, _ := cr.Marshal(requestMessage)

		provider.On("SendApplicationMessageToNode", mock.Anything, expectedIEEE, expectedAppMessage, false).Return(nil).Run(func(args mock.Arguments) {
			message := zcl.Message{
				FrameType:           zcl.FrameGlobal,
				Direction:           zcl.ServerToClient,
				TransactionSequence: 255,
				Manufacturer:        0,
				ClusterID:           0x0001,
				SourceEndpoint:      2,
				DestinationEndpoint: 1,
				Command: &global.ReadAttributesResponse{
					Records: []global.ReadAttributeResponseRecord{},
				},
			}

			appMessageReply, _ := cr.Marshal(message)

			c.ProcessIncomingMessage(zigbee.NodeIncomingMessageEvent{
				Node: zigbee.Node{
					IEEEAddress:    expectedIEEE,
					NetworkAddress: 0,
					LogicalType:    0,
					LQI:            0,
					Depth:          0,
					LastDiscovered: time.Time{},
					LastReceived:   time.Time{},
				},
				IncomingMessage: zigbee.IncomingMessage{
					GroupID:              0,
					SourceIEEEAddress:    expectedIEEE,
					SourceNetworkAddress: 0,
					Broadcast:            false,
					Secure:               false,
					LinkQuality:          0,
					Sequence:             0,
					ApplicationMessage:   appMessageReply,
				},
			})
		})

		responseMessage, err := c.RequestResponse(context.Background(), expectedIEEE, false, requestMessage)
		assert.NoError(t, err)

		assert.Equal(t, requestMessage.TransactionSequence, responseMessage.TransactionSequence)
		assert.IsType(t, &global.ReadAttributesResponse{}, responseMessage.Command)

		provider.AssertExpectations(t)
	})

	t.Run("a message via the communicator goes to provider and a reply which does not match address is not returned to the caller", func(t *testing.T) {
		provider := &zigbee.MockProvider{}
		cr := zcl.NewCommandRegistry()
		global.Register(cr)

		c := NewCommunicator(provider, cr)

		expectedIEEE := zigbee.IEEEAddress(2)
		requestMessage := zcl.Message{
			FrameType:           zcl.FrameGlobal,
			Direction:           zcl.ClientToServer,
			TransactionSequence: 255,
			Manufacturer:        0,
			ClusterID:           0x0001,
			SourceEndpoint:      1,
			DestinationEndpoint: 2,
			Command: &global.ReadAttributes{
				Identifier: []zcl.AttributeID{1},
			},
		}

		expectedAppMessage, _ := cr.Marshal(requestMessage)

		provider.On("SendApplicationMessageToNode", mock.Anything, expectedIEEE, expectedAppMessage, false).Return(nil).Run(func(args mock.Arguments) {
			message := zcl.Message{
				FrameType:           zcl.FrameGlobal,
				Direction:           zcl.ServerToClient,
				TransactionSequence: 255,
				Manufacturer:        0,
				ClusterID:           0x0001,
				SourceEndpoint:      2,
				DestinationEndpoint: 1,
				Command: &global.ReadAttributesResponse{
					Records: []global.ReadAttributeResponseRecord{
						{
							Identifier: 1,
							Status:     0,
							DataTypeValue: &zcl.AttributeDataTypeValue{
								DataType: zcl.TypeSignedInt8,
								Value:    int64(64),
							},
						},
					},
				},
			}

			appMessageReply, _ := cr.Marshal(message)

			c.ProcessIncomingMessage(zigbee.NodeIncomingMessageEvent{
				Node: zigbee.Node{
					IEEEAddress:    zigbee.IEEEAddress(3),
					NetworkAddress: 0,
					LogicalType:    0,
					LQI:            0,
					Depth:          0,
					LastDiscovered: time.Time{},
					LastReceived:   time.Time{},
				},
				IncomingMessage: zigbee.IncomingMessage{
					GroupID:              0,
					SourceIEEEAddress:    expectedIEEE,
					SourceNetworkAddress: 0,
					Broadcast:            false,
					Secure:               false,
					LinkQuality:          0,
					Sequence:             0,
					ApplicationMessage:   appMessageReply,
				},
			})
		})

		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		_, err := c.RequestResponse(ctx, expectedIEEE, false, requestMessage)
		assert.Error(t, err)

		provider.AssertExpectations(t)
	})

	t.Run("a message via the communicator goes to provider and a reply which does not match the sequence is not returned to the caller", func(t *testing.T) {
		provider := &zigbee.MockProvider{}
		cr := zcl.NewCommandRegistry()
		global.Register(cr)

		c := NewCommunicator(provider, cr)

		expectedIEEE := zigbee.IEEEAddress(2)
		requestMessage := zcl.Message{
			FrameType:           zcl.FrameGlobal,
			Direction:           zcl.ClientToServer,
			TransactionSequence: 255,
			Manufacturer:        0,
			ClusterID:           0x0001,
			SourceEndpoint:      1,
			DestinationEndpoint: 2,
			Command: &global.ReadAttributes{
				Identifier: []zcl.AttributeID{1},
			},
		}

		expectedAppMessage, _ := cr.Marshal(requestMessage)

		provider.On("SendApplicationMessageToNode", mock.Anything, expectedIEEE, expectedAppMessage, false).Return(nil).Run(func(args mock.Arguments) {
			message := zcl.Message{
				FrameType:           zcl.FrameGlobal,
				Direction:           zcl.ServerToClient,
				TransactionSequence: 1,
				Manufacturer:        0,
				ClusterID:           0x0001,
				SourceEndpoint:      2,
				DestinationEndpoint: 1,
				Command: &global.ReadAttributesResponse{
					Records: []global.ReadAttributeResponseRecord{
						{
							Identifier: 1,
							Status:     0,
							DataTypeValue: &zcl.AttributeDataTypeValue{
								DataType: zcl.TypeSignedInt8,
								Value:    int64(64),
							},
						},
					},
				},
			}

			appMessageReply, _ := cr.Marshal(message)

			c.ProcessIncomingMessage(zigbee.NodeIncomingMessageEvent{
				Node: zigbee.Node{
					IEEEAddress:    expectedIEEE,
					NetworkAddress: 0,
					LogicalType:    0,
					LQI:            0,
					Depth:          0,
					LastDiscovered: time.Time{},
					LastReceived:   time.Time{},
				},
				IncomingMessage: zigbee.IncomingMessage{
					GroupID:              0,
					SourceIEEEAddress:    expectedIEEE,
					SourceNetworkAddress: 0,
					Broadcast:            false,
					Secure:               false,
					LinkQuality:          0,
					Sequence:             0,
					ApplicationMessage:   appMessageReply,
				},
			})
		})

		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		_, err := c.RequestResponse(ctx, expectedIEEE, false, requestMessage)
		assert.Error(t, err)

		provider.AssertExpectations(t)
	})
}

func TestCommunicator_GlobalReadAttributes(t *testing.T) {
	t.Run("requests attributes from provider and returns responses", func(t *testing.T) {
		mockProvider := &zigbee.MockProvider{}
		cr := zcl.NewCommandRegistry()
		global.Register(cr)

		ieee := zigbee.IEEEAddress(0x0102030405060708)

		c := NewCommunicator(mockProvider, cr)
		g := c.Global()

		expectedValue := "value"
		clusterId := zigbee.ClusterID(0x1223)

		srcEndpoint := zigbee.Endpoint(4)
		destEndpoint := zigbee.Endpoint(8)

		transactionSequence := uint8(0x7f)

		expectedRequestOne := zcl.Message{
			FrameType:           zcl.FrameGlobal,
			Direction:           zcl.ClientToServer,
			TransactionSequence: transactionSequence,
			Manufacturer:        zigbee.NoManufacturer,
			ClusterID:           clusterId,
			SourceEndpoint:      srcEndpoint,
			DestinationEndpoint: destEndpoint,
			Command: &global.ReadAttributes{
				Identifier: []zcl.AttributeID{0x0004, 0x0005},
			},
		}

		expectedResponse := &global.ReadAttributesResponse{
			Records: []global.ReadAttributeResponseRecord{
				{
					Identifier: 0x0004,
					Status:     0,
					DataTypeValue: &zcl.AttributeDataTypeValue{
						DataType: zcl.TypeStringCharacter8,
						Value:    expectedValue,
					},
				},
				{
					Identifier: 0x0005,
					Status:     1,
				},
			},
		}

		appRequestOne, _ := cr.Marshal(expectedRequestOne)

		mockProvider.On("SendApplicationMessageToNode", mock.Anything, ieee, appRequestOne, true).Return(nil).Run(func(args mock.Arguments) {
			message := zcl.Message{
				FrameType:           zcl.FrameGlobal,
				Direction:           zcl.ServerToClient,
				TransactionSequence: transactionSequence,
				Manufacturer:        zigbee.NoManufacturer,
				ClusterID:           clusterId,
				SourceEndpoint:      destEndpoint,
				DestinationEndpoint: srcEndpoint,
				Command:             expectedResponse,
			}

			appMessageReply, _ := cr.Marshal(message)

			c.ProcessIncomingMessage(zigbee.NodeIncomingMessageEvent{
				Node: zigbee.Node{
					IEEEAddress:    ieee,
					NetworkAddress: 0,
					LogicalType:    0,
					LQI:            0,
					Depth:          0,
					LastDiscovered: time.Time{},
					LastReceived:   time.Time{},
				},
				IncomingMessage: zigbee.IncomingMessage{
					GroupID:              0,
					SourceIEEEAddress:    ieee,
					SourceNetworkAddress: 0,
					Broadcast:            false,
					Secure:               false,
					LinkQuality:          0,
					Sequence:             0,
					ApplicationMessage:   appMessageReply,
				},
			})
		})

		resp, err := g.ReadAttributes(context.Background(), ieee, true, clusterId, zigbee.NoManufacturer, srcEndpoint, destEndpoint, transactionSequence, []zcl.AttributeID{0x0004, 0x0005})
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse.Records, resp)
	})
}

func TestCommunicator_GlobalConfigureReporting(t *testing.T) {
	t.Run("requests attributes from provider and returns responses", func(t *testing.T) {
		mockProvider := &zigbee.MockProvider{}
		cr := zcl.NewCommandRegistry()
		global.Register(cr)

		ieee := zigbee.IEEEAddress(0x0102030405060708)

		c := NewCommunicator(mockProvider, cr)
		g := c.Global()

		clusterId := zigbee.ClusterID(0x1223)

		srcEndpoint := zigbee.Endpoint(4)
		destEndpoint := zigbee.Endpoint(8)

		transactionSequence := uint8(0x7f)

		attributeId := zcl.AttributeID(0x0001)
		dataType := zcl.TypeUnsignedInt8
		minInterval := uint16(0x0000)
		maxInterval := uint16(0xffff)
		reportableChange := uint64(10)

		expectedRequestOne := zcl.Message{
			FrameType:           zcl.FrameGlobal,
			Direction:           zcl.ClientToServer,
			TransactionSequence: transactionSequence,
			Manufacturer:        zigbee.NoManufacturer,
			ClusterID:           clusterId,
			SourceEndpoint:      srcEndpoint,
			DestinationEndpoint: destEndpoint,
			Command: &global.ConfigureReporting{
				Records: []global.ConfigureReportingRecord{
					{
						Direction:        0,
						Identifier:       attributeId,
						DataType:         dataType,
						MinimumInterval:  minInterval,
						MaximumInterval:  maxInterval,
						ReportableChange: &zcl.AttributeDataValue{Value: reportableChange},
						Timeout:          0,
					},
				},
			},
		}

		expectedResponse := &global.ConfigureReportingResponse{
			Records: []global.ConfigureReportingResponseRecord{
				{
					Status:     0,
					Direction:  0,
					Identifier: attributeId,
				},
			},
		}

		appRequestOne, _ := cr.Marshal(expectedRequestOne)

		mockProvider.On("SendApplicationMessageToNode", mock.Anything, ieee, appRequestOne, true).Return(nil).Run(func(args mock.Arguments) {
			message := zcl.Message{
				FrameType:           zcl.FrameGlobal,
				Direction:           zcl.ServerToClient,
				TransactionSequence: transactionSequence,
				Manufacturer:        zigbee.NoManufacturer,
				ClusterID:           clusterId,
				SourceEndpoint:      destEndpoint,
				DestinationEndpoint: srcEndpoint,
				Command:             expectedResponse,
			}

			appMessageReply, _ := cr.Marshal(message)

			c.ProcessIncomingMessage(zigbee.NodeIncomingMessageEvent{
				Node: zigbee.Node{
					IEEEAddress:    ieee,
					NetworkAddress: 0,
					LogicalType:    0,
					LQI:            0,
					Depth:          0,
					LastDiscovered: time.Time{},
					LastReceived:   time.Time{},
				},
				IncomingMessage: zigbee.IncomingMessage{
					GroupID:              0,
					SourceIEEEAddress:    ieee,
					SourceNetworkAddress: 0,
					Broadcast:            false,
					Secure:               false,
					LinkQuality:          0,
					Sequence:             0,
					ApplicationMessage:   appMessageReply,
				},
			})
		})

		err := g.ConfigureReporting(context.Background(), ieee, true, clusterId, zigbee.NoManufacturer, srcEndpoint, destEndpoint, transactionSequence, attributeId, dataType, minInterval, maxInterval, reportableChange)
		assert.NoError(t, err)
	})
}
