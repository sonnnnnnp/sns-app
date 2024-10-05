package user

import (
	"context"

	"github.com/google/uuid"
	user_repository "github.com/sonnnnnnp/sns-app/internal/domain/user"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IUserUsecase interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetUserByUsername(ctx context.Context, username string) (*ent.User, *oapi.SocialContext, error)
	GetUserFollowers(ctx context.Context, id uuid.UUID) ([]*ent.User, error)
	GetUserFollowing(ctx context.Context, id uuid.UUID) ([]*ent.User, error)
	FollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialContext, error)
	UnfollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialContext, error)
	CreateUser(ctx context.Context) (*ent.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, body *oapi.UpdateUserJSONBody) (*oapi.User, error)
}

type UserUsecase struct {
	userRepo *user_repository.UserRepository
}

func New(
	userRepo *user_repository.UserRepository,
) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

var _ IUserUsecase = (*UserUsecase)(nil)
