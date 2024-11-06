package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetUserByName(ctx context.Context, name string) (*db.User, error) {
	queries := db.New(repo.pool)

	u, err := queries.GetUserByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrUserNotFound
		}
		return nil, err
	}

	return &u, nil
}
