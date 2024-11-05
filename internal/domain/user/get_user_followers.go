package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (repo *UserRepository) GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]*ent.User, error) {
	// return ur.db.User.Query().Where(user.ID(uID)).QueryFollowers().All(ctx)

	return nil, nil
}
