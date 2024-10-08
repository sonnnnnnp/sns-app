package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IPostRepository interface {
	GetPosts(ctx context.Context, limit *int, fromCursor *uuid.UUID, uID *uuid.UUID) (posts []*ent.Post, nextCursor uuid.UUID, err error)
	CreatePost(ctx context.Context, uID uuid.UUID, body *oapi.CreatePostJSONBody) (*ent.Post, error)
	GetPostFavorited(ctx context.Context, uID uuid.UUID, pID uuid.UUID) (bool, error)
	GetPostFavoritesCount(ctx context.Context, pID uuid.UUID) (int, error)
	FavoritePost(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error
	UnfavoritePost(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error
}

type PostRepository struct {
	db *ent.Client
}

func New(db *ent.Client) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

var _ IPostRepository = (*PostRepository)(nil)
