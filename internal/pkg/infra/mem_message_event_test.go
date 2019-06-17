package infra_test

import (
	"testing"
	"time"

	"github.com/wxdao/chatroom/internal/pkg/domain"
	"github.com/wxdao/chatroom/internal/pkg/infra"
)

func TestMemoryMessageEventChannel(t *testing.T) {
	ec := infra.NewMemoryMessageEventChannel()

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
