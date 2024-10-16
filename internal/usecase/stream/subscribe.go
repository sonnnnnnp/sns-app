package stream

import (
	"context"
	"fmt"

	"github.com/sonnnnnnp/sns-app/internal/tools/ws"
)

func (su *StreamUsecase) OnSubscribe(ctx context.Context, hub *ws.Hub, client *ws.Client, msg ws.Message) {
	if err := client.Subscribe(msg.Channel); err != nil {
		client.Send(ws.Message{
			Type: "error",
			Body: fmt.Sprintf("%v", err),
		})
		return
	}

	client.Send(ws.Message{
		Type: "subscribe",
		Body: fmt.Sprintf("subscribed to channel %s", msg.Channel),
	})
}

func (su *StreamUsecase) OnUnsubscribe(ctx context.Context, h *ws.Hub, client *ws.Client, msg ws.Message) {
	if err := client.Unsubscribe(msg.Channel); err != nil {
		client.Send(ws.Message{
			Type: "error",
			Body: fmt.Sprintf("%v", err),
		})
		return
	}

	client.Send(ws.Message{
		Type: "unsubscribe",
		Body: fmt.Sprintf("unsubscribed from channel %s", msg.Channel),
	})
}
