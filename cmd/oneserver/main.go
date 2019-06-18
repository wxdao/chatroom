package main

import (
	"context"
	"flag"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/wxdao/chatroom/internal/pkg/api"
	"github.com/wxdao/chatroom/internal/pkg/infra"
	"github.com/wxdao/chatroom/pkg/message"
	"github.com/wxdao/chatroom/pkg/push"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var messageAddr = flag.String("message_addr", "0.0.0.0:8000", "Listen address of message service")
var pushAddr = flag.String("push_addr", "0.0.0.0:8001", "Listen address of push service")
var pprofAddr = flag.String("pprof_addr", "0.0.0.0:6000", "Listen address of http pprof")
var configFile = flag.String("config", "config.yaml", "Path to the config file")

type syncReader struct {
	mux    sync.Mutex
	reader io.Reader
}

func (sr *syncReader) Read(p []byte) (n int, err error) {
	sr.mux.Lock()
	defer sr.mux.Unlock()

	return sr.reader.Read(p)
}

func main() {
	flag.Parse()

	conf := parseConfigFromFile(*configFile)

	errc := make(chan error)

	// start http pprof
	go func() {
		log.Println("started http pprof on", *pprofAddr)
		errc <- http.ListenAndServe(*pprofAddr, nil)
	}()

	// setup mongo client
	mgoClient, err := mongo.NewClient(
		options.Client().ApplyURI(conf.MongoURI),
	)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mgoClient.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	messageColl := mgoClient.Database(conf.MessageDB).Collection(conf.MessageColl)

	// setup entropy
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	entropy := &syncReader{reader: rand.New(source)}

	// setup message event channel
	var eventChannel api.MessageEventChannel
	if conf.NatsAddr != "" {
		nc, err := nats.Connect(conf.NatsAddr)
		if err != nil {
			log.Fatalln(err)
		}
		eventChannel = infra.NewNatsMessageEventChannel(nc, conf.NatsSubjectPrefix)
	} else {
		eventChannel = infra.NewMemoryMessageEventChannel()
	}

	// start message service server
	messageLis, err := net.Listen("tcp", *messageAddr)
	if err != nil {
		log.Fatalln(err)
	}
	messageGRPCServer := grpc.NewServer()
	message.RegisterMessageServiceServer(messageGRPCServer, &api.DomainMessageServiceServer{
		MessageRepository:   infra.NewMgoMessageRepository(messageColl),
		Entropy:             entropy,
		MessageEventChannel: eventChannel,
	})
	go func() {
		log.Println("started message service server on", *messageAddr)
		errc <- messageGRPCServer.Serve(messageLis)
	}()

	// start push service server
	pushLis, err := net.Listen("tcp", *pushAddr)
	if err != nil {
		log.Fatalln(err)
	}
	pushGRPCServer := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime: 30 * time.Second,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time: time.Minute,
		}),
	)
	push.RegisterPushServiceServer(pushGRPCServer, &api.DomainPushServiceServer{
		MessageEventChannel: eventChannel,
	})
	go func() {
		log.Println("started push service server on", *pushAddr)
		errc <- pushGRPCServer.Serve(pushLis)
	}()

	log.Fatal(<-errc)
}
