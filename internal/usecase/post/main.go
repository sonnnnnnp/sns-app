package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/domain/post"
	"github.com/sonnnnnnp/sns-app/internal/domain/user"
	"github.com/sonnnnnnp/sns-app/internal/usecase/stream"
)

type IPostUsecase interface {
	// posts
	CreatePost(ctx context.Context, body *api.CreatePostJSONBody) (*api.Post, error)

	GetPostByID(ctx context.Context, pID uuid.UUID) (*api.Post, error)

	DeletePost(ctx context.Context, pID uuid.UUID) error

	// favorites
	CreatePostFavorite(ctx context.Context, pID uuid.UUID) error

	GetPostFavorites(ctx context.Context, pID uuid.UUID) ([]api.PostFavorite, error)

	DeletePostFavorite(ctx context.Context, pID uuid.UUID) error

	// timeline
	GetTimeline(ctx context.Context, params *api.GetTimelineParams) (posts []api.Post, nextCursor uuid.UUID, err error)
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
