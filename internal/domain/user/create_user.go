package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) CreateUser(ctx context.Context, lineID *string) (*db.User, error) {
	queries := db.New(repo.conn)

	nickname := "unknown"
	u, err := queries.CreateUser(ctx, db.CreateUserParams{
		ID:       uuid.New(),
		Name:     uuid.NewString(),
		Nickname: &nickname,
		LineID:   lineID,
	})
	if err != nil {
		return nil, err
	}

	return &u, nil
}
