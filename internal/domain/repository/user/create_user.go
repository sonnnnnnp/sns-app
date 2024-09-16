package user_repository

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
)

func (ur *UserRepository) CreateUser(ctx context.Context, id string) (*ent.User, error) {
	u, err := ur.db.User.Create().SetUsername(id).SetDisplayName(id).Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
