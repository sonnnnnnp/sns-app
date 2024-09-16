package usecase

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	repository "github.com/sonnnnnnp/sns-app/internal/domain/repository/user"
)

type IUserUsecase interface {
	CreateUser(ctx context.Context, id string) (*ent.User, error)
}

type UserUsecase struct {
	userRepo *repository.UserRepository
}

func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

var _ IUserUsecase = (*UserUsecase)(nil)
