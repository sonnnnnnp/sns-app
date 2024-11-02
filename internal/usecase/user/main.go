package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/user"
	"github.com/sonnnnnnp/sns-app/pkg/db"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IUserUsecase interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*oapi.User, error)
	GetUserByName(ctx context.Context, name string) (*oapi.User, error)
	GetUserFollowers(ctx context.Context, id uuid.UUID) ([]oapi.User, error)
	GetUserFollowing(ctx context.Context, id uuid.UUID) ([]oapi.User, error)
	FollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialContext, error)
	UnfollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialContext, error)
	RemoveUserFromFollowers(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialContext, error)
	CreateUser(ctx context.Context) (*db.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, body *oapi.UpdateUserJSONBody) (*oapi.User, error)
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
