package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/user"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IUserUsecase interface {
	// users
	GetUserByID(ctx context.Context, uID uuid.UUID) (*oapi.User, error)
	GetUserByName(ctx context.Context, name string) (*oapi.User, error)

	UpdateUser(ctx context.Context, body *oapi.UpdateUserJSONBody) (*oapi.User, error)

	// blocks
	BlockUser(ctx context.Context, uID uuid.UUID) error

	GetUserBlocking(ctx context.Context) ([]oapi.User, error)

	UnblockUser(ctx context.Context, uID uuid.UUID) error

	// follows
	FollowUser(ctx context.Context, uID uuid.UUID) (*oapi.SocialConnection, error)

	GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]oapi.UserFollower, error)
	GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]oapi.UserFollower, error)

	UnfollowUser(ctx context.Context, uID uuid.UUID) (*oapi.SocialConnection, error)
	RemoveUserFromFollowers(ctx context.Context, uID uuid.UUID) (*oapi.SocialConnection, error)
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
