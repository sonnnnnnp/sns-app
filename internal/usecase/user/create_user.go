package usecase

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
)

func (uu *UserUsecase) CreateUser(ctx context.Context, id string) (*ent.User, error) {
	u, err := uu.userRepo.CreateUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
