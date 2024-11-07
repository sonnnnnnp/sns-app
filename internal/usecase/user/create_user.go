package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *UserUsecase) CreateUser(ctx context.Context) (*db.User, error) {
	return uc.userRepo.CreateUser(ctx, nil)
}
