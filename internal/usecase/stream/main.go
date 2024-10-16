package stream

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ws"
)

type IStreamUsecase interface {
	OnSubscribe(ctx context.Context, hub *ws.Hub, client *ws.Client, msg ws.Message)
	OnUnsubscribe(ctx context.Context, hub *ws.Hub, client *ws.Client, msg ws.Message)
}

type StreamUsecase struct{}

func New() *StreamUsecase {
	return &StreamUsecase{}
}

var _ IStreamUsecase = (*StreamUsecase)(nil)
