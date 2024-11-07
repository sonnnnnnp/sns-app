package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*db.User, error) {
	queries := db.New(repo.pool)

	u, err := queries.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
