package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/post"
	"github.com/sonnnnnnp/sns-app/internal/domain/user"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IPostUsecase interface {
	CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*ent.Post, error)
	FavoritePost(ctx context.Context, pID uuid.UUID) error
	UnavoritePost(ctx context.Context, pID uuid.UUID) error
}

type PostUsecase struct {
	postRepo *post.PostRepository
	userRepo *user.UserRepository
}

func New(
	postRepo *post.PostRepository,
	userRepo *user.UserRepository,
) *PostUsecase {
	return &PostUsecase{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

var _ IPostUsecase = (*PostUsecase)(nil)
