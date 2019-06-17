package api

import (
	context "context"
	"io"
	"log"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/wxdao/chatroom/internal/pkg/domain"
	"github.com/wxdao/chatroom/pkg/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DomainMessageServiceServer is an implement of rpc MessageServiceServer using domain logic
type DomainMessageServiceServer struct {
	MessageRepository   domain.MessageRepository
	Entropy             io.Reader
	MessageEventChannel MessageEventChannel
}

// QueryMessagesInRange implements the rpc method QueryMessagesInRange
func (s *DomainMessageServiceServer) QueryMessagesInRange(
	ctx context.Context,
	req *message.QueryMessagesInRangeRequest,
) (*message.QueryMessagesInRangeReply, error) {
	msgs, err := s.MessageRepository.MessageInRange(req.ChatroomID, req.StartID, req.EndID)
	if err != nil {
		log.Println("ERROR|QueryMessagesInRange|MessageRepository.MessageInRange", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	var reply message.QueryMessagesInRangeReply
	for _, msg := range msgs {

		reply.Messages = append(reply.Messages, &message.Message{
			Id:         msg.ID(),
			ChatroomID: msg.ChatroomID(),
			Username:   msg.Username(),
			Content:    msg.Content(),
			CreateTime: msg.CreateTime().UnixNano(),
		})
	}
	return &reply, nil
}

// SendMessage implements the rpc method SendMessage
func (s *DomainMessageServiceServer) SendMessage(
	ctx context.Context,
	req *message.SendMessageRequest,
) (*message.SendMessageReply, error) {
	currentTime := time.Now()
	id, err := ulid.New(ulid.Timestamp(currentTime), s.Entropy)
	if err != nil {
		log.Println("ERROR|SendMessage|ulid.New", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	msg := domain.NewMessage(id.String(), req.ChatroomID, req.Username, req.Content, currentTime)
	err = s.MessageRepository.SaveMessage(msg)
	if err != nil {
		log.Println("ERROR|SendMessage|MessageRepository.SaveMessage", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = s.MessageEventChannel.EmitNewMessageEvent(msg)
	if err != nil {
		log.Println("ERROR|SendMessage|MessageEventChannel.EmitNewMessageEvent", err)
	}
	return &message.SendMessageReply{
		Message: &message.Message{
			Id:         msg.ID(),
			ChatroomID: msg.ChatroomID(),
			Username:   msg.Username(),
			Content:    msg.Content(),
			CreateTime: msg.CreateTime().UnixNano(),
		},
	}, nil
}
