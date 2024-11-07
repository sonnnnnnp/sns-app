package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) UnfollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error {
	queries := db.New(repo.pool)

	if err := queries.DeleteUserFollower(ctx, db.DeleteUserFollowerParams{
		FollowerID:  selfUID,
		FollowingID: targetUID,
	}); err != nil {
		return err
	}

	return nil
}
