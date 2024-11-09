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
	// posts
	CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*oapi.Post, error)

	GetPostByID(ctx context.Context, pID uuid.UUID) (*oapi.Post, error)

	DeletePost(ctx context.Context, pID uuid.UUID) error

	// favorites
	CreatePostFavorite(ctx context.Context, pID uuid.UUID) error

	GetPostFavorites(ctx context.Context, pID uuid.UUID) ([]oapi.PostFavorite, error)

	DeletePostFavorite(ctx context.Context, pID uuid.UUID) error

	// timeline
	GetTimeline(ctx context.Context, params *oapi.GetTimelineParams) (posts []oapi.Post, nextCursor uuid.UUID, err error)
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
