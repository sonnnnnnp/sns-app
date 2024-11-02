package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (ur *UserRepository) GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]*ent.User, error) {
	// return ur.db.User.Query().Where(user.ID(uID)).QueryFollowing().All(ctx)

	return nil, nil
}
