package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) BlockUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error {
	return db.New(repo.pool).CreateUserBlock(ctx, db.CreateUserBlockParams{
		BlockerID:  selfUID,
		BlockingID: targetUID,
	})
}
