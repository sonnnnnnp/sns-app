package user

import (
	"context"

	"github.com/google/uuid"
)

func (repo *UserRepository) UnfollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error {
	// return ur.db.User.UpdateOneID(selfUID).RemoveFollowingIDs(targetUID).Exec(ctx)

	return nil
}
