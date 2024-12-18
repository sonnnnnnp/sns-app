package gateway

import (
	"context"

	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
)

type IGatewayUsecase interface {
	GetGatewayURI(ctx context.Context) (*api.GatewayURI, error)
}

type GatewayUsecase struct{}

func New() *GatewayUsecase {
	return &GatewayUsecase{}
}

var _ IGatewayUsecase = (*GatewayUsecase)(nil)
