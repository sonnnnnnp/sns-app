package call_timeline

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/reverie/internal/server/http/adapter/api"
)

type ICallTimelineUsecase interface {
	GetCallTimeline(ctx context.Context, params api.GetCallTimelineParams) (rooms []api.CallRoom, nextCursor uuid.UUID, err error)
	GetFollowingCallTimeline(ctx context.Context, params api.GetFollowingCallTimelineParams) (rooms []api.CallRoom, nextCursor uuid.UUID, err error)
	GetUserCallTimeline(ctx context.Context, uID uuid.UUID, params api.GetUserCallTimelineParams) (rooms []api.CallRoom, nextCursor uuid.UUID, err error)
}

type CallTimelineUsecase struct {
	pool *pgxpool.Pool
}

func New(
	pool *pgxpool.Pool,
) *CallTimelineUsecase {
	return &CallTimelineUsecase{
		pool: pool,
	}
}

var (
	defaultLimit = 25
	maxLimit     = 100
)

var _ ICallTimelineUsecase = (*CallTimelineUsecase)(nil)
