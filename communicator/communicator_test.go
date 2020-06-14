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

	t.Run("a parsable message can be retrieved by ReadMessage", func(t *testing.T) {
		provider := &zigbee.MockProvider{}
		cr := zcl.NewCommandRegistry()
		global.Register(cr)

		c := NewCommunicator(provider, cr)

		expectedIEEE := zigbee.IEEEAddress(1)

		expectedMessage := zcl.Message{
			FrameType:           zcl.FrameGlobal,
			Direction:           zcl.ServerToClient,
			TransactionSequence: 255,
			Manufacturer:        0,
			ClusterID:           0x0001,
			SourceEndpoint:      1,
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

		appMessage, err := cr.Marshal(expectedMessage)
		assert.NoError(t, err)

		err = c.ProcessIncomingMessage(zigbee.NodeIncomingMessageEvent{
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
				LinkQuality:          1,
				Sequence:             1,
				ApplicationMessage:   appMessage,
			},
		})

		assert.NoError(t, err)

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		messageWithSource, err := c.ReadMessage(ctx)

		assert.NoError(t, err)
		assert.Equal(t, expectedIEEE, messageWithSource.SourceAddress)
		assert.Equal(t, expectedMessage, messageWithSource.Message)

		assert.Equal(t, 0, len(c.matches))
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

		provider.On("SendNodeMessageToNode", mock.Anything, expectedIEEE, expectedAppMessage).Return(nil)

		err := c.Request(context.Background(), expectedIEEE, message)
		assert.NoError(t, err)

		provider.AssertExpectations(t)
	})
}

func TestCommunicator_RequestResponse(t *testing.T) {
	t.Run("a message via the communicator goes to provider and the reply is returned to the caller", func(t *testing.T) {
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

		provider.On("SendNodeMessageToNode", mock.Anything, expectedIEEE, expectedAppMessage).Return(nil).Run(func(args mock.Arguments) {
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

		responseMessage, err := c.RequestResponse(context.Background(), expectedIEEE, requestMessage)
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

		provider.On("SendNodeMessageToNode", mock.Anything, expectedIEEE, expectedAppMessage).Return(nil).Run(func(args mock.Arguments) {
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

		_, err := c.RequestResponse(ctx, expectedIEEE, requestMessage)
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

		provider.On("SendNodeMessageToNode", mock.Anything, expectedIEEE, expectedAppMessage).Return(nil).Run(func(args mock.Arguments) {
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

		_, err := c.RequestResponse(ctx, expectedIEEE, requestMessage)
		assert.Error(t, err)

		provider.AssertExpectations(t)
	})
}
