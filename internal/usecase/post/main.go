package post

import (
	"context"

	post_repository "github.com/sonnnnnnp/sns-app/internal/domain/post"
	user_repository "github.com/sonnnnnnp/sns-app/internal/domain/user"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IPostUsecase interface {
	CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*ent.Post, error)
}

type PostUsecase struct {
	postRepo *post_repository.PostRepository
	userRepo *user_repository.UserRepository
}

func New(
	postRepo *post_repository.PostRepository,
	userRepo *user_repository.UserRepository,
) *PostUsecase {
	return &PostUsecase{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

var _ IPostUsecase = (*PostUsecase)(nil)
