package call

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
)

type ICallUsecase interface {
	CreateCall(ctx context.Context, body api.CreateCallJSONBody) (*api.CallRoom, error)
	EndCall(ctx context.Context, cID uuid.UUID) error
}

type CallUsecase struct {
	pool *pgxpool.Pool
}

func New(
	pool *pgxpool.Pool,
) *CallUsecase {
	return &CallUsecase{
		pool: pool,
	}
}

var _ ICallUsecase = (*CallUsecase)(nil)
