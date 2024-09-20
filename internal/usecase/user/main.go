package user_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	user_repository "github.com/sonnnnnnp/sns-app/internal/domain/repository/user"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IUserUsecase interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
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
