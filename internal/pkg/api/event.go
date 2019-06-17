package api

import "github.com/wxdao/chatroom/internal/pkg/domain"

// MessageEventChannel defines the interface of the message event channel
type MessageEventChannel interface {
	EmitNewMessageEvent(message *domain.Message) error
	Subscribe(chatroomID string, handler func(m *domain.Message)) (drain func(), err error)
	Drain() error
}
