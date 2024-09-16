package user_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
)

func (uu *UserUsecase) GetUser(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	u, err := uu.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
