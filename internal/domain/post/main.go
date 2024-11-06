package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/pkg/db"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IPostRepository interface {
	GetPosts(ctx context.Context, limit *int, fromCursor *uuid.UUID, uID *uuid.UUID) (rows []db.GetPostsRow, nextCursor uuid.UUID, err error)
	CreatePost(ctx context.Context, uID uuid.UUID, body *oapi.CreatePostJSONBody) (*db.Post, error)
	GetPostFavorited(ctx context.Context, uID uuid.UUID, pID uuid.UUID) (bool, error)
	GetPostFavoritesCount(ctx context.Context, pID uuid.UUID) (int, error)
	FavoritePost(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error
	UnfavoritePost(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error
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
