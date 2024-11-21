package timeline

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
)

type ITimelineUsecase interface {
	GetTimeline(ctx context.Context, params *api.GetTimelineParams) (posts []api.Post, nextCursor uuid.UUID, err error)
	GetFollowingTimeline(ctx context.Context, params *api.GetFollowingTimelineParams) (posts []api.Post, nextCursor uuid.UUID, err error)
	GetUserTimeline(ctx context.Context, uID uuid.UUID, params *api.GetUserTimelineParams) (posts []api.Post, nextCursor uuid.UUID, err error)
}

type TimelineUsecase struct {
	pool *pgxpool.Pool
}

func New(
	pool *pgxpool.Pool,
) *TimelineUsecase {
	return &TimelineUsecase{
		pool: pool,
	}
}

var (
	defaultLimit = 25
	maxLimit     = 100
)

var _ ITimelineUsecase = (*TimelineUsecase)(nil)
