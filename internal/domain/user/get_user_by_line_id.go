package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetUserByLineID(ctx context.Context, lineID string) (*db.User, error) {
	queries := db.New(repo.pool)

	u, err := queries.GetUserByLineID(ctx, &lineID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
