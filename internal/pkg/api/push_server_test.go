package api_test

import (
	"testing"
	"time"

	"github.com/wxdao/chatroom/internal/pkg/api"
	"github.com/wxdao/chatroom/internal/pkg/domain"
	"github.com/wxdao/chatroom/internal/pkg/infra"
	"github.com/wxdao/chatroom/pkg/push"
	"google.golang.org/grpc"
)

type messagesServer struct {
	grpc.ServerStream

	c chan *push.MessagesReply
}

func (s *messagesServer) Send(reply *push.MessagesReply) error {
	s.c <- reply
	return nil
}

func TestDomainPushServiceServerMessages(t *testing.T) {
	ec := infra.NewMemoryMessageEventChannel()
	server := api.DomainPushServiceServer{
		MessageEventChannel: ec,
	}

	ms := &messagesServer{
		c: make(chan *push.MessagesReply),
	}
	go server.Messages(&push.MessagesRequest{
		ChatroomID: "c1",
	}, ms)

	// wait for server to be done subscribing
	time.Sleep(10 * time.Millisecond)

	err := ec.EmitNewMessageEvent(domain.NewMessage("1", "c1", "u1", "content1", time.Now()))
	if err != nil {
		t.Fatal(err)
	}

	select {
	case reply := <-ms.c:
		if reply.Message.Content != "content1" {
			t.Fatal("wrong content")
		}
	case <-time.NewTimer(time.Second).C:
		t.Fatal("timeout")
	}
}
