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
