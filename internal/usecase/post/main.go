package post

import (
	"context"

	post_repository "github.com/sonnnnnnp/sns-app/internal/domain/post"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IPostUsecase interface {
	CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*ent.Post, error)
}

type PostUsecase struct {
	postRepo *post_repository.PostRepository
}

func New(
	postRepo *post_repository.PostRepository,
) *PostUsecase {
	return &PostUsecase{
		postRepo: postRepo,
	}
}

var _ IPostUsecase = (*PostUsecase)(nil)
