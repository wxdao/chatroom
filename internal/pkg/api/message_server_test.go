package api_test

import (
	context "context"
	"testing"
	"time"

	"github.com/wxdao/chatroom/internal/pkg/api"
	"github.com/wxdao/chatroom/internal/pkg/infra"
	"github.com/wxdao/chatroom/pkg/message"
)

func TestDomainMessageServiceServerQueryMessagesInRange(t *testing.T) {
	cleanMessageColl(t)
	server := api.DomainMessageServiceServer{
		MessageRepository:   infra.NewMgoMessageRepository(messageColl),
		Entropy:             entropy,
		MessageEventChannel: infra.NewMemoryMessageEventChannel(),
	}

	var ids []string
	for i := 0; i < 5; i++ {
		reply, err := server.SendMessage(context.TODO(), &message.SendMessageRequest{
			ChatroomID: "testchatroom",
			Username:   "testusername",
			Content:    "testcontent",
		})
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, reply.Message.Id)
		time.Sleep(10 * time.Millisecond)
	}

	reply, err := server.QueryMessagesInRange(context.TODO(), &message.QueryMessagesInRangeRequest{
		ChatroomID: "testchatroom",
		StartID:    ids[1],
		EndID:      "a", // get all to the latest
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(reply.Messages) != 4 {
		t.Fatal("should get 4 messages, got", len(reply.Messages))
	}

	reply2, err := server.QueryMessagesInRange(context.TODO(), &message.QueryMessagesInRangeRequest{
		ChatroomID: "testchatroom",
		StartID:    "0", // get all from the oldest
		EndID:      ids[3],
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(reply2.Messages) != 4 {
		t.Fatal("should get 4 messages, got", len(reply2.Messages))
	}
}

func TestDomainMessageServiceServerSendMessage(t *testing.T) {
	cleanMessageColl(t)
	server := api.DomainMessageServiceServer{
		MessageRepository:   infra.NewMgoMessageRepository(messageColl),
		Entropy:             entropy,
		MessageEventChannel: infra.NewMemoryMessageEventChannel(),
	}

	reply, err := server.SendMessage(context.TODO(), &message.SendMessageRequest{
		ChatroomID: "testchatroom",
		Username:   "testusername",
		Content:    "testcontent",
	})
	if err != nil {
		t.Fatal(err)
	}
	if reply.Message.ChatroomID != "testchatroom" {
		t.Fatal("wrong ChatroomID")
	}
	if reply.Message.Username != "testusername" {
		t.Fatal("wrong Username")
	}
	if reply.Message.Content != "testcontent" {
		t.Fatal("wrong Content")
	}
}
