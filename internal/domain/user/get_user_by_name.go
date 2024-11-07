package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetUserByName(ctx context.Context, name string) (*db.User, error) {
	queries := db.New(repo.pool)

	u, err := queries.GetUserByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
