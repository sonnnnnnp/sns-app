package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) UnblockUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error {
	return db.New(repo.pool).DeleteUserBlock(ctx, db.DeleteUserBlockParams{
		BlockerID:  selfUID,
		BlockingID: targetUID,
	})
}
