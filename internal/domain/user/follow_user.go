package user

import (
	"context"

	"github.com/google/uuid"
)

func (repo *UserRepository) FollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error {
	// return ur.db.User.UpdateOneID(selfUID).AddFollowingIDs(targetUID).Exec(ctx)

	return nil
}
