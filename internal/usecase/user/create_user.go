package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uu *UserUsecase) CreateUser(ctx context.Context) (*db.User, error) {
	return uu.userRepo.CreateUser(ctx, nil)
}
