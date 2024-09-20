package user_usecase

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
)

func (uu *UserUsecase) CreateUser(ctx context.Context) (*ent.User, error) {
	u, err := uu.userRepo.CreateUser(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
