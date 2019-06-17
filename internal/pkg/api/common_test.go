package api_test

import (
	context "context"
	"io"
	"log"
	"math/rand"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var messageColl *mongo.Collection

var entropy io.Reader

func init() {
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
	messageColl = mgoClient.Database("apitest").Collection("message")

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
