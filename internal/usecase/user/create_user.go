package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (uu *UserUsecase) CreateUser(ctx context.Context) (*ent.User, error) {
	return uu.userRepo.CreateUser(ctx)
}
