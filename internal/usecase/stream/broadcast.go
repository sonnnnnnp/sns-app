package stream

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/internal/tools/ws"
)

// if uIDs is nil, broadcasts to all users with the given channel
func (su *StreamUsecase) Broadcast(ctx context.Context, channel string, msg ws.Message, uIDs []uuid.UUID) {
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
