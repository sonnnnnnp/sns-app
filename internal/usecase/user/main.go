package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/user"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IUserUsecase interface {
	// users
	GetUserByID(ctx context.Context, id uuid.UUID) (*oapi.User, error)
	GetUserByName(ctx context.Context, name string) (*oapi.User, error)

	UpdateUser(ctx context.Context, id uuid.UUID, body *oapi.UpdateUserJSONBody) (*oapi.User, error)

	// followers
	FollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialConnection, error)

	GetUserFollowers(ctx context.Context, id uuid.UUID) ([]oapi.UserFollower, error)
	GetUserFollowing(ctx context.Context, id uuid.UUID) ([]oapi.UserFollower, error)

	UnfollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialConnection, error)
	RemoveUserFromFollowers(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialConnection, error)
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
