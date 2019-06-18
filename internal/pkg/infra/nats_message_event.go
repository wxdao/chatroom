package infra

import (
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/wxdao/chatroom/internal/pkg/domain"
)

type natsMessage struct {
	ID         string `json:"id"`
	ChatroomID string `json:"chatroomID"`
	Username   string `json:"username"`
	Content    string `json:"content"`
	CreateTime int64  `json:"createTime"`
}

// NatsMessageEventChannel is an implement of api.MessageEventChannel using nats
type NatsMessageEventChannel struct {
	nec    *nats.EncodedConn
	prefix string

	mux      sync.Mutex
	lastID   int
	handlers map[string][]handlerObj
	subs     map[string]*nats.Subscription
}

// NewNatsMessageEventChannel creates a new NatsMessageEventChannel
func NewNatsMessageEventChannel(nc *nats.Conn, subjectPrefix string) *NatsMessageEventChannel {
	nec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	return &NatsMessageEventChannel{
		nec:      nec,
		prefix:   subjectPrefix,
		handlers: map[string][]handlerObj{},
		subs:     map[string]*nats.Subscription{},
	}
}

func (ec *NatsMessageEventChannel) subject(chatroomID string) string {
	return ec.prefix + ".chatroom." + chatroomID
}

// EmitNewMessageEvent implements EmitNewMessageEvent
func (ec *NatsMessageEventChannel) EmitNewMessageEvent(message *domain.Message) error {
	nm := natsMessage{
		ID:         message.ID(),
		ChatroomID: message.ChatroomID(),
		Username:   message.Username(),
		Content:    message.Content(),
		CreateTime: message.CreateTime().UnixNano(),
	}
	return ec.nec.Publish(ec.subject(message.ChatroomID()), nm)
}

// Subscribe implements Subscribe
func (ec *NatsMessageEventChannel) Subscribe(chatroomID string, handler func(m *domain.Message)) (func(), error) {
	ec.mux.Lock()
	defer ec.mux.Unlock()

	ec.lastID++
	ho := handlerObj{
		id:      ec.lastID,
		handler: handler,
	}
	ec.handlers[chatroomID] = append(ec.handlers[chatroomID], ho)

	if ec.subs[chatroomID] == nil {
		sub, err := ec.nec.Subscribe(ec.subject(chatroomID), func(nm *natsMessage) {
			ec.mux.Lock()
			defer ec.mux.Unlock()

			handlers := ec.handlers[nm.ChatroomID]
			for _, ho := range handlers {
				go ho.handler(&domain.Message{
					Xid:         nm.ID,
					XchatroomID: nm.ChatroomID,
					Xusername:   nm.Username,
					Xcontent:    nm.Content,
					XcreateTime: time.Unix(0, nm.CreateTime),
				})
			}
		})
		if err != nil {
			return nil, err
		}
		ec.subs[chatroomID] = sub
	}

	return func() {
		ec.mux.Lock()
		defer ec.mux.Unlock()

		for i, v := range ec.handlers[chatroomID] {
			if v.id == ho.id {
				ec.handlers[chatroomID] = append(ec.handlers[chatroomID][:i], ec.handlers[chatroomID][i+1:]...)
				if len(ec.handlers[chatroomID]) == 0 && ec.subs[chatroomID] != nil {
					ec.subs[chatroomID].Unsubscribe()
					delete(ec.subs, chatroomID)
				}
				return
			}
		}
	}, nil
}

// Drain implements Drain
func (ec *NatsMessageEventChannel) Drain() error {
	return nil
}
