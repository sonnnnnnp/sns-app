package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/usecase/stream"
)

type IPostUsecase interface {
	// posts
	CreatePost(ctx context.Context, body *api.CreatePostJSONBody) (*api.Post, error)

	GetPostByID(ctx context.Context, pID uuid.UUID) (*api.Post, error)

	DeletePost(ctx context.Context, pID uuid.UUID) error

	// favorites
	CreatePostFavorite(ctx context.Context, pID uuid.UUID) error

	GetPostFavorites(ctx context.Context, pID uuid.UUID) ([]api.PostFavorite, error)

	DeletePostFavorite(ctx context.Context, pID uuid.UUID) error

	// timeline
	GetTimeline(ctx context.Context, params *api.GetTimelineParams) (posts []api.Post, nextCursor uuid.UUID, err error)
}

type PostUsecase struct {
	pool *pgxpool.Pool

	streamUsecase *stream.StreamUsecase
}

func New(
	pool *pgxpool.Pool,
	streamUsecase *stream.StreamUsecase,
) *PostUsecase {
	return &PostUsecase{
		pool: pool,

		streamUsecase: streamUsecase,
	}
}

var _ IPostUsecase = (*PostUsecase)(nil)
