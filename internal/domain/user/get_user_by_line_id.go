package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetUserByLineID(ctx context.Context, lineID string) (*db.User, error) {
	queries := db.New(repo.pool)

	u, err := queries.GetUserByLineID(ctx, &lineID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrUserNotFound
		}
		return nil, err
	}

	return &u, nil
}
