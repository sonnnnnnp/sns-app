package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (ur *UserRepository) CreateUser(ctx context.Context, lineID *string) (*ent.User, error) {
	u, err := ur.db.User.Create().SetName(uuid.New().String()).SetLineID(*lineID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
