package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]db.GetUserFollowersRow, error) {
	queries := db.New(repo.pool)

	rows, err := queries.GetUserFollowers(ctx, uID)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
