package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]db.GetUserFollowingRow, error) {
	queries := db.New(repo.pool)

	rows, err := queries.GetUserFollowing(ctx, uID)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
