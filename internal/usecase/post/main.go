package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/post"
	"github.com/sonnnnnnp/sns-app/internal/domain/user"
	"github.com/sonnnnnnp/sns-app/internal/usecase/stream"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IPostUsecase interface {
	CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*oapi.Post, error)
	CreatePostFavorite(ctx context.Context, pID uuid.UUID) error
	DeletePostFavorite(ctx context.Context, pID uuid.UUID) error
	GetPostFavorites(ctx context.Context, pID uuid.UUID) (favorites []oapi.PostFavorite, err error)
}

type PostUsecase struct {
	postRepo *post.PostRepository
	userRepo *user.UserRepository

	streamUsecase *stream.StreamUsecase
}

func New(
	postRepo *post.PostRepository,
	userRepo *user.UserRepository,
	streamUsecase *stream.StreamUsecase,
) *PostUsecase {
	return &PostUsecase{
		postRepo: postRepo,
		userRepo: userRepo,

		streamUsecase: streamUsecase,
	}
}

var _ IPostUsecase = (*PostUsecase)(nil)
