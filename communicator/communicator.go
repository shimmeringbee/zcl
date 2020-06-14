package communicator

import (
	"context"
	"errors"
	"fmt"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zigbee"
	"sync"
	"sync/atomic"
)

type MessageWithSource struct {
	SourceAddress zigbee.IEEEAddress
	Message       zcl.Message
}

type Matcher func(address zigbee.IEEEAddress, appMsg zigbee.ApplicationMessage, zclMessage zcl.Message) bool

func AddressAndSequenceMatch(matchAddress zigbee.IEEEAddress, matchSequence uint8) Matcher {
	return func(address zigbee.IEEEAddress, appMsg zigbee.ApplicationMessage, zclMessage zcl.Message) bool {
		return matchAddress == address && matchSequence == zclMessage.TransactionSequence
	}
}

func (c *Communicator) NewMatch(matcher Matcher, callback func(source MessageWithSource)) Match {
	return Match{
		Id:       atomic.AddUint64(c.matchId, 1),
		Matcher:  matcher,
		Callback: callback,
	}
}

type Match struct {
	Id       uint64
	Matcher  Matcher
	Callback func(source MessageWithSource)
}

type Communicator struct {
	Provider        zigbee.Provider
	CommandRegistry *zcl.CommandRegistry

	readableMessages chan MessageWithSource

	mutex   *sync.RWMutex
	matches map[uint64]Match
	matchId *uint64
}

const DefaultMessagesToBeRead = 50

func NewCommunicator(provider zigbee.Provider, registry *zcl.CommandRegistry) *Communicator {
	return &Communicator{
		Provider:         provider,
		CommandRegistry:  registry,
		readableMessages: make(chan MessageWithSource, DefaultMessagesToBeRead),
		mutex:            &sync.RWMutex{},
		matches:          map[uint64]Match{},
		matchId:          new(uint64),
	}
}

func (c *Communicator) ProcessIncomingMessage(msg zigbee.NodeIncomingMessageEvent) error {
	message, err := c.CommandRegistry.Unmarshal(msg.ApplicationMessage)

	if err != nil {
		return fmt.Errorf("failed to unmarshal incomming ZCL message: %w", err)
	}

	c.mutex.RLock()
	ourMatches := c.matches
	c.mutex.RUnlock()

	for _, match := range ourMatches {
		if match.Matcher(msg.IEEEAddress, msg.ApplicationMessage, message) {
			go match.Callback(MessageWithSource{
				SourceAddress: msg.IEEEAddress,
				Message:       message,
			})
		}
	}

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

	c.matches[match.Id] = match
}

func (c *Communicator) removeMatch(match Match) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.matches, match.Id)
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

	match := c.NewMatch(AddressAndSequenceMatch(address, message.TransactionSequence),
		func(recvMessage MessageWithSource) {
			ch <- recvMessage.Message
		})

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
