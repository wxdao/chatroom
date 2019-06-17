package infra_test

import (
	"context"
	"testing"
	"time"

	"github.com/wxdao/chatroom/internal/pkg/domain"
	"github.com/wxdao/chatroom/internal/pkg/infra"
)

func TestMemoryMessageEventChannel(t *testing.T) {
	ec := infra.NewMemoryMessageEventChannel()

	c := make(chan string)
	err := ec.Subscribe("c1", func(m *domain.Message) {
		c <- m.Content()
	})
	if err != nil {
		t.Fatal(err)
	}

	ec.EmitNewMessageEvent(domain.NewMessage("1", "c2", "u1", "testcontent", time.Now()))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	select {
	case <-c:
		t.Fatal("should not have received events")
	case <-ctx.Done():
	}

	ec.EmitNewMessageEvent(domain.NewMessage("1", "c1", "u1", "testcontent", time.Now()))

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	select {
	case content := <-c:
		if content != "testcontent" {
			t.Fatal("wrong content")
		}
	case <-ctx.Done():
		t.Fatal("timeout")
	}
}
