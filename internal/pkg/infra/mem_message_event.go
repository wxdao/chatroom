package infra

import (
	"sync"

	"github.com/wxdao/chatroom/internal/pkg/domain"
)

type handlerObj struct {
	id      int
	handler func(m *domain.Message)
}

// MemoryMessageEventChannel is an in-memory implement of api.MessageEventChannel
type MemoryMessageEventChannel struct {
	mux      sync.Mutex
	lastID   int
	handlers map[string][]handlerObj
}

// NewMemoryMessageEventChannel creates a new MemoryMessageEventChannel
func NewMemoryMessageEventChannel() *MemoryMessageEventChannel {
	return &MemoryMessageEventChannel{
		handlers: map[string][]handlerObj{},
	}
}

// EmitNewMessageEvent implements EmitNewMessageEvent
func (ec *MemoryMessageEventChannel) EmitNewMessageEvent(message *domain.Message) error {
	ec.mux.Lock()
	defer ec.mux.Unlock()

	handlers := ec.handlers[message.ChatroomID()]
	for _, ho := range handlers {
		go ho.handler(message)
	}
	return nil
}

// Subscribe implements Subscribe
func (ec *MemoryMessageEventChannel) Subscribe(chatroomID string, handler func(m *domain.Message)) (func(), error) {
	ec.mux.Lock()
	defer ec.mux.Unlock()

	ec.lastID++
	ho := handlerObj{
		id:      ec.lastID,
		handler: handler,
	}
	ec.handlers[chatroomID] = append(ec.handlers[chatroomID], ho)

	return func() {
		ec.mux.Lock()
		defer ec.mux.Unlock()

		for i, v := range ec.handlers[chatroomID] {
			if v.id == ho.id {
				ec.handlers[chatroomID] = append(ec.handlers[chatroomID][:i], ec.handlers[chatroomID][i+1:]...)
				return
			}
		}
	}, nil
}

// Drain implements Drain
func (ec *MemoryMessageEventChannel) Drain() error {
	return nil
}
