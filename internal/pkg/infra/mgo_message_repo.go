package infra

import (
	"context"
	"time"

	"github.com/wxdao/chatroom/internal/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MgoMessageRepository is an implementation of domain.MessageRepository using mongodb
type MgoMessageRepository struct {
	messageColl *mongo.Collection
}

// NewMgoMessageRepository creates a new MgoMessageRepository
func NewMgoMessageRepository(messageColl *mongo.Collection) *MgoMessageRepository {
	return &MgoMessageRepository{messageColl: messageColl}
}

// MessageByID implements MessageRepository.MessageByID
func (repo *MgoMessageRepository) MessageByID(id string) (*domain.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	var model mgoMessageModel
	err := repo.messageColl.FindOne(ctx, filter).Decode(&model)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return model.message(), nil
}

// MessageInRange implements MessageRepository.MessageInRange
func (repo *MgoMessageRepository) MessageInRange(startID string, endID string) ([]*domain.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"id": bson.M{"$lte": endID, "$gte": startID}}
	cur, err := repo.messageColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var messages []*domain.Message
	for cur.Next(ctx) {
		var model mgoMessageModel
		err := cur.Decode(&model)
		if err != nil {
			return nil, err
		}
		messages = append(messages, model.message())
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

// SaveMessage implements MessageRepository.SaveMessage
func (repo *MgoMessageRepository) SaveMessage(message *domain.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	model := newMgoMessageModel(message)
	_, err := repo.messageColl.InsertOne(ctx, model)
	if err != nil {
		return err
	}
	return nil
}

type mgoMessageModel struct {
	ID         string `bson:"id"`
	ChatroomID string `bson:"chatroomID"`
	Username   string `bson:"username"`
	Content    string `bson:"content"`
}

func newMgoMessageModel(message *domain.Message) *mgoMessageModel {
	var m mgoMessageModel
	m.ID = message.Xid
	m.ChatroomID = message.XchatroomID
	m.Username = message.Xusername
	m.Content = message.Xcontent
	return &m
}

func (m mgoMessageModel) message() *domain.Message {
	var dm domain.Message
	dm.Xid = m.ID
	dm.XchatroomID = m.ChatroomID
	dm.Xusername = m.Username
	dm.Xcontent = m.Content
	return &dm
}
