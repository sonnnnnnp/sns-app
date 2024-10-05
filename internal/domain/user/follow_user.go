package user

import (
	"context"

	"github.com/google/uuid"
)

func (ur *UserRepository) FollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error {
	return ur.db.User.UpdateOneID(selfUID).AddFollowingIDs(targetUID).Exec(ctx)
}
