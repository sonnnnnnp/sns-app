package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) FollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error {
	queries := db.New(repo.pool)

	if err := queries.CreateUserFollower(ctx, db.CreateUserFollowerParams{
		FollowerID:  targetUID,
		FollowingID: selfUID,
	}); err != nil {
		return err
	}

	return nil
}
