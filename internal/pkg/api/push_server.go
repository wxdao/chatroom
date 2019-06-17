package api

import (
	"log"

	"github.com/wxdao/chatroom/internal/pkg/domain"
	"github.com/wxdao/chatroom/pkg/push"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DomainPushServiceServer is an implement of rpc PushServiceServer using domain logic
type DomainPushServiceServer struct {
	MessageEventChannel MessageEventChannel
}

// Messages implements rpc method Messages
func (s *DomainPushServiceServer) Messages(req *push.MessagesRequest, stream push.PushService_MessagesServer) error {
	c := make(chan error)
	defer close(c)
	drain, err := s.MessageEventChannel.Subscribe(req.ChatroomID, func(m *domain.Message) {
		err := stream.Send(&push.MessagesReply{
			Message: &push.Message{
				Id:         m.ID(),
				ChatroomID: m.ChatroomID(),
				Username:   m.Username(),
				Content:    m.Content(),
				CreateTime: m.CreateTime().UnixNano(),
			},
		})
		if err != nil {
			c <- err
		}
	})
	if err != nil {
		log.Println("ERROR|Messages|MessageEventChannel.Subscribe", err)
		return status.Error(codes.Internal, err.Error())
	}
	defer drain()
	err = <-c
	if err != nil {
		log.Println("ERROR|Messages|BEFORE_RETURN", err)
	}
	return err
}
