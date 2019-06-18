package infra_test

import (
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/wxdao/chatroom/internal/pkg/domain"
	"github.com/wxdao/chatroom/internal/pkg/infra"
)

func TestNatsMessageEventChannel(t *testing.T) {
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Drain()
	ec := infra.NewNatsMessageEventChannel(nc, "test")

	c := make(chan string)
	drain, err := ec.Subscribe("c1", func(m *domain.Message) {
		c <- m.Content()
	})
	if err != nil {
		t.Fatal(err)
	}

	ec.EmitNewMessageEvent(domain.NewMessage("1", "c2", "u1", "testcontent", time.Now()))
	select {
	case <-c:
		t.Fatal("should not have received events")
	case <-time.NewTimer(time.Second).C:
	}

	ec.EmitNewMessageEvent(domain.NewMessage("1", "c1", "u1", "testcontent", time.Now()))
	select {
	case content := <-c:
		if content != "testcontent" {
			t.Fatal("wrong content")
		}
	case <-time.NewTimer(time.Second).C:
		t.Fatal("timeout")
	}

	drain()
	ec.EmitNewMessageEvent(domain.NewMessage("1", "c1", "u1", "testcontent", time.Now()))
	select {
	case <-c:
		t.Fatal("should not have received events")
	case <-time.NewTimer(time.Second).C:
	}
}
