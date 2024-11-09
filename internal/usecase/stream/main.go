package stream

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ws"
)

type IStreamUsecase interface {
	OnSubscribe(ctx context.Context, client *ws.Client, msg ws.Message) error
	OnUnsubscribe(ctx context.Context, client *ws.Client, msg ws.Message) error

	Broadcast(ctx context.Context, channel string, msg ws.Message, uIDs []uuid.UUID)
}

type StreamUsecase struct{}

func New() *StreamUsecase {
	return &StreamUsecase{}
}

var _ IStreamUsecase = (*StreamUsecase)(nil)
