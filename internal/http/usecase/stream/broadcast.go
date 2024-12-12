package stream

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/reverie/internal/http/tools/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/ws"
)

// if uIDs is nil, broadcasts to all users with the given channel
func (uc *StreamUsecase) Broadcast(ctx context.Context, channel string, msg ws.Message, uIDs []uuid.UUID) {
	hub := ctxhelper.GetWebSocketHub(ctx)

	msg.Channel = channel
	for client := range hub.Clients {
		if client.Channels[channel] {
			if uIDs == nil || hub.ContainsUser(uIDs, client.UserID) {
				client.Send(msg)
			}
		}
	}
}
