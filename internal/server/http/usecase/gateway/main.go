package gateway

import (
	"context"

	"github.com/sonnnnnnp/reverie/internal/server/http/adapter/api"
)

type IGatewayUsecase interface {
	GetGatewayURI(ctx context.Context) (*api.GatewayURI, error)
}

type GatewayUsecase struct{}

func New() *GatewayUsecase {
	return &GatewayUsecase{}
}

var _ IGatewayUsecase = (*GatewayUsecase)(nil)
