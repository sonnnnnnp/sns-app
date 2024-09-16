package user_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
)

func (ur *UserRepository) GetUser(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	u, err := ur.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
