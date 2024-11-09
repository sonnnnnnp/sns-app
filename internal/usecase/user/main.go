package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/domain/user"
)

type IUserUsecase interface {
	// users
	GetUserByID(ctx context.Context, uID uuid.UUID) (*api.User, error)
	GetUserByName(ctx context.Context, name string) (*api.User, error)

	UpdateUser(ctx context.Context, body *api.UpdateUserJSONBody) (*api.User, error)

	// blocks
	BlockUser(ctx context.Context, uID uuid.UUID) error

	GetUserBlocking(ctx context.Context) ([]api.User, error)

	UnblockUser(ctx context.Context, uID uuid.UUID) error

	// follows
	FollowUser(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error)

	GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]api.UserFollower, error)
	GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]api.UserFollower, error)

	UnfollowUser(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error)
	RemoveUserFromFollowers(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error)
}

type UserUsecase struct {
	userRepo *user.UserRepository
}

func New(
	userRepo *user.UserRepository,
) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

var _ IUserUsecase = (*UserUsecase)(nil)
