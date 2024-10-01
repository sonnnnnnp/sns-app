package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (uu *UserUsecase) GetUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	u, err := uu.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return u, nil
}
