package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wxdao/chatroom/pkg/message"
	"github.com/wxdao/chatroom/pkg/push"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var messageURL string
var pushURL string
var chatroomID string
var username string

var messageClient message.MessageServiceClient
var pushClient push.PushServiceClient

var app *tview.Application
var history *tview.List
var input *tview.InputField

var messages = map[string]*msg{}
var messagesMux sync.Mutex

type msg struct {
	id         string
	username   string
	content    string
	createTime time.Time
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: client <message service url> <push service url> <chatroom> <username>")
		return
	}
	messageURL = os.Args[1]
	pushURL = os.Args[2]
	chatroomID = os.Args[3]
	username = os.Args[4]

	setupClients()

	setupUI()

	go loopForMessages()

	input.SetDoneFunc(func(key tcell.Key) {
		if input.GetText() == "" {
			return
		}
		text := input.GetText()
		go func() {
			_, err := messageClient.SendMessage(context.Background(), &message.SendMessageRequest{
				ChatroomID: chatroomID,
				Username:   username,
				Content:    text,
			})
			if err != nil {
				log.Fatal("failed to call message.SendMessage:", err)
			}
		}()
		input.SetText("")
	})

	if err := app.Run(); err != nil {
		log.Fatal("failed to run app:", err)
	}
}

func setupClients() {
	messageConn, err := grpc.Dial(messageURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to dial message service:", err)
	}
	messageClient = message.NewMessageServiceClient(messageConn)

	pushConn, err := grpc.Dial(pushURL, grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time: time.Minute,
	}))
	if err != nil {
		log.Fatal("failed to dial push service:", err)
	}
	pushClient = push.NewPushServiceClient(pushConn)
}

func setupUI() {
	history = tview.NewList()
	history.SetSelectedFocusOnly(true)
	history.ShowSecondaryText(false).SetBorder(true).SetTitle(fmt.Sprintf(" CHATROOM(%s) ", chatroomID))

	input = tview.NewInputField()
	input.SetLabel("> ")

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(history, 0, 9, false).
		AddItem(input, 3, 1, true)

	app = tview.NewApplication().SetRoot(flex, true)

	historyFocus := false
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTAB {
			historyFocus = !historyFocus
			if historyFocus {
				app.SetFocus(history)
			} else {
				app.SetFocus(input)
			}
			return nil
		}
		return event
	})
}

func loopForMessages() {
	for {
		mc, err := pushClient.Messages(context.Background(), &push.MessagesRequest{
			ChatroomID: chatroomID,
		})
		if err != nil {
			log.Fatal("failed to call push.Messages:", err)
		}
		// query for potencial missing messages
		queryLatestMessages()
		// loop recving latest messages
		for {
			reply, err := mc.Recv()
			if err != nil {
				log.Println(err)
				break
			}
			messagesMux.Lock()
			messages[reply.Message.Id] = &msg{
				id:         reply.Message.Id,
				username:   reply.Message.Username,
				content:    reply.Message.Content,
				createTime: time.Unix(0, reply.Message.CreateTime),
			}
			messagesMux.Unlock()
			app.QueueUpdateDraw(func() {
				updateHistory()
			})
		}
	}
}

func queryLatestMessages() {
	keys := sortedMessagesKeys()
	startID := "0"
	if len(keys) != 0 {
		startID = keys[len(keys)-1]
	}
	reply, err := messageClient.QueryMessagesInRange(context.Background(), &message.QueryMessagesInRangeRequest{
		ChatroomID: chatroomID,
		StartID:    startID,
		EndID:      "a",
	})
	if err != nil {
		log.Fatal("failed to call message.QueryMessagesInRange:", err)
	}
	messagesMux.Lock()
	for _, m := range reply.Messages {
		messages[m.Id] = &msg{
			id:         m.Id,
			username:   m.Username,
			content:    m.Content,
			createTime: time.Unix(0, m.CreateTime),
		}
	}
	messagesMux.Unlock()
	app.QueueUpdateDraw(func() {
		updateHistory()
	})
}

func updateHistory() {
	messagesMux.Lock()
	defer messagesMux.Unlock()

	keys := sortedMessagesKeys()

	history.Clear()
	for _, k := range keys {
		msg := messages[k]
		history.AddItem(
			fmt.Sprintf(
				"%s <%s>: %s",
				msg.createTime.Format("2006-01-02 15:04:05 MST"),
				msg.username,
				msg.content,
			),
			"", 0, nil)
	}
	history.SetCurrentItem(-1)
}

func sortedMessagesKeys() (keys []string) {
	for k := range messages {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
