package post_usecase

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	post_repository "github.com/sonnnnnnp/sns-app/internal/domain/repository/post"
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
