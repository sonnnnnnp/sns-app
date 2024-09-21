package post_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IPostRepository interface {
	CreatePost(ctx context.Context, uID uuid.UUID, body *oapi.CreatePostJSONBody) (*ent.Post, *ent.User, error)
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
