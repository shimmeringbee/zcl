package communicator

import (
	"context"
	"errors"
	"fmt"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
	"sync"
)

type MessageWithSource struct {
	SourceAddress zigbee.IEEEAddress
	Message       zcl.Message
}

type Match struct {
	Address  zigbee.IEEEAddress
	Sequence uint8
	Fn       func(zcl.Message)
}

type Communicator struct {
	Provider        zigbee.Provider
	CommandRegistry *zcl.CommandRegistry

	readableMessages chan MessageWithSource

	mutex   *sync.RWMutex
	matches []Match
}

const DefaultMessagesToBeRead = 50

func NewCommunicator(provider zigbee.Provider, registry *zcl.CommandRegistry) *Communicator {
	return &Communicator{
		Provider:         provider,
		CommandRegistry:  registry,
		readableMessages: make(chan MessageWithSource, DefaultMessagesToBeRead),
		mutex:            &sync.RWMutex{},
	}
}

func (c *Communicator) ProcessIncomingMessage(msg zigbee.NodeIncomingMessageEvent) error {
	message, err := c.CommandRegistry.Unmarshal(msg.ApplicationMessage)

	if err != nil {
		return fmt.Errorf("failed to unmarshal incomming ZCL message: %w", err)
	}

	c.mutex.RLock()
	for _, match := range c.matches {
		if match.Address == msg.IEEEAddress && match.Sequence == message.TransactionSequence {
			go match.Fn(message)
		}
	}
	c.mutex.RUnlock()

	select {
	case c.readableMessages <- MessageWithSource{
		SourceAddress: msg.IEEEAddress,
		Message:       message,
	}:
	default:
		return fmt.Errorf("ZCL communicator readable message channel is full")
	}

	return nil
}

func (c *Communicator) ReadMessage(ctx context.Context) (MessageWithSource, error) {
	select {
	case message := <-c.readableMessages:
		return message, nil
	case <-ctx.Done():
		return MessageWithSource{}, fmt.Errorf("ZCL communicator read message context expired")
	}
}

func (c *Communicator) addMatch(match Match) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.matches = append(c.matches, match)
}

func (c *Communicator) removeMatch(match Match) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	originalMatches := c.matches
	c.matches = []Match{}

	for _, iMatch := range originalMatches {
		if match.Address != iMatch.Address || match.Sequence != iMatch.Sequence {
			c.matches = append(c.matches, iMatch)
		}
	}
}

func (c *Communicator) Request(ctx context.Context, address zigbee.IEEEAddress, message zcl.Message) error {
	appMessage, err := c.CommandRegistry.Marshal(message)

	if err != nil {
		return fmt.Errorf("ZCL communicator failed to send message during marshalling: %w", err)
	}

	err = c.Provider.SendNodeMessageToNode(ctx, address, appMessage)

	if err != nil {
		return fmt.Errorf("ZCL communicator failed to send via provider: %w", err)
	}

	return nil
}

func (c *Communicator) RequestResponse(ctx context.Context, address zigbee.IEEEAddress, message zcl.Message) (zcl.Message, error) {
	ch := make(chan zcl.Message, 1)

	match := Match{
		Address:  address,
		Sequence: message.TransactionSequence,
		Fn: func(recvMessage zcl.Message) {
			ch <- recvMessage
		},
	}

	c.addMatch(match)
	defer c.removeMatch(match)

	if err := c.Request(ctx, address, message); err != nil {
		return zcl.Message{}, err
	}

	select {
	case resp := <-ch:
		return resp, nil
	case <-ctx.Done():
		return zcl.Message{}, errors.New("ZCL communicator waiting for reply, context expired")
	}
}
