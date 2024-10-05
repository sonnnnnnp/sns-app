package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (uu *UserUsecase) GetUserFollowing(ctx context.Context, id uuid.UUID) ([]*ent.User, error) {
	u, err := uu.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return uu.userRepo.GetUserFollowing(ctx, u)
}
