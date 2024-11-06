package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) CreateUser(ctx context.Context, lineID *string) (*db.User, error) {
	queries := db.New(repo.pool)

	u, err := queries.CreateUser(ctx, lineID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
