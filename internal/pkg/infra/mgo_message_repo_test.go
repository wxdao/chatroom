package infra_test

import (
	"context"
	"io"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/wxdao/chatroom/internal/pkg/domain"
	"github.com/wxdao/chatroom/internal/pkg/infra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var messageColl *mongo.Collection

var entropy io.Reader

func init() {
	var err error
	mgoClient, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = mgoClient.Connect(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	messageColl = mgoClient.Database("test").Collection("message")

	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	entropy = rand.New(source)
}

func cleanMessageColl(t *testing.T) {
	_, err := messageColl.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMgoMessageRepositoryMessageByID(t *testing.T) {
	cleanMessageColl(t)
	repo := infra.NewMgoMessageRepository(messageColl)

	id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	msg, err := repo.MessageByID(id)
	if err != nil {
		t.Fatal(err)
	}
	if msg != nil {
		t.Fatal("msg should be nil")
	}

	err = repo.SaveMessage(
		domain.NewMessage(id, "testchatroom", "testusername", "testcontent"),
	)
	if err != nil {
		t.Fatal(err)
	}
	msg, err = repo.MessageByID(id)
	if err != nil {
		t.Fatal(err)
	}
	if msg == nil {
		t.Fatal("msg should not be nil")
	}
}

func TestMgoMessageRepositoryMessageInRange(t *testing.T) {
	cleanMessageColl(t)
	repo := infra.NewMgoMessageRepository(messageColl)

	ids := []string{
		ulid.MustNew(ulid.Timestamp(time.Unix(1, 0)), entropy).String(),
		ulid.MustNew(ulid.Timestamp(time.Unix(2, 0)), entropy).String(),
		ulid.MustNew(ulid.Timestamp(time.Unix(3, 0)), entropy).String(),
		ulid.MustNew(ulid.Timestamp(time.Unix(4, 0)), entropy).String(),
		ulid.MustNew(ulid.Timestamp(time.Unix(5, 0)), entropy).String(),
	}

	for i, id := range ids {
		err := repo.SaveMessage(
			domain.NewMessage(id, "testchatroom", "testusername", strconv.Itoa(i)),
		)
		if err != nil {
			t.Fatal(err)
		}
	}
	err := repo.SaveMessage(
		domain.NewMessage(ids[2], "testchatroom2", "testusername", strconv.Itoa(2)),
	)
	if err != nil {
		t.Fatal(err)
	}

	msgs, err := repo.MessageInRange("testchatroom", ids[1], ids[3])
	if err != nil {
		t.Fatal(err)
	}
	if len(msgs) != 3 {
		t.Fatal("should get 3 messages, got", len(msgs))
	}
	for i, msg := range msgs {
		x, _ := strconv.Atoi(msg.Content())
		if x != i+1 {
			t.Fatal("bad order")
		}
		if msg.ChatroomID() != "testchatroom" {
			t.Fatal("unexpected chatroom id")
		}
	}
}
