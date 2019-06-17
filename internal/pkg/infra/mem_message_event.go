package infra

import (
	"sync"

	"github.com/wxdao/chatroom/internal/pkg/domain"
)

// MemoryMessageEventChannel is an in-memory implement of api.MessageEventChannel
type MemoryMessageEventChannel struct {
	mux      sync.Mutex
	handlers map[string][]func(m *domain.Message)
}

// NewMemoryMessageEventChannel creates a new MemoryMessageEventChannel
func NewMemoryMessageEventChannel() *MemoryMessageEventChannel {
	return &MemoryMessageEventChannel{
		handlers: map[string][]func(m *domain.Message){},
	}
}

// EmitNewMessageEvent implements EmitNewMessageEvent
func (ec *MemoryMessageEventChannel) EmitNewMessageEvent(message *domain.Message) error {
	ec.mux.Lock()
	defer ec.mux.Unlock()

	handlers := ec.handlers[message.ChatroomID()]
	for _, handler := range handlers {
		go handler(message)
	}
	return nil
}

// Subscribe implements Subscribe
func (ec *MemoryMessageEventChannel) Subscribe(chatroomID string, handler func(m *domain.Message)) error {
	ec.mux.Lock()
	defer ec.mux.Unlock()

	ec.handlers[chatroomID] = append(ec.handlers[chatroomID], handler)
	return nil
}

// Drain implements Drain
func (ec *MemoryMessageEventChannel) Drain() error {
	return nil
}
