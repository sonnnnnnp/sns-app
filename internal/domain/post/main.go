package post

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

type IPostRepository interface {
	// posts
	CreatePost(ctx context.Context, uID uuid.UUID, body *api.CreatePostJSONBody) (pID uuid.UUID, err error)

	GetPostByID(ctx context.Context, pID uuid.UUID) (*db.GetPostByIDRow, error)

	DeletePost(ctx context.Context, pID uuid.UUID) error

	// favorites
	CreatePostFavorite(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error

	GetPostFavorite(ctx context.Context, uID uuid.UUID, pID uuid.UUID) (*db.GetPostFavoriteRow, error)
	GetPostFavorites(ctx context.Context, pID uuid.UUID) ([]db.GetPostFavoritesRow, error)

	DeletePostFavorite(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error

	// timeline
	GetTimeline(ctx context.Context, limit *int, fromCursor *time.Time, targetUID *uuid.UUID, following *bool) (rows []db.GetTimelineRow, err error)
}

type PostRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *PostRepository {
	return &PostRepository{
		pool: pool,
	}
}

var _ IPostRepository = (*PostRepository)(nil)
