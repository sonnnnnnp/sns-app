package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetBlockStatus(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) (db.GetBlockStatusRow, error) {
	return db.New(repo.pool).GetBlockStatus(ctx, db.GetBlockStatusParams{
		SelfID:   selfUID,
		TargetID: targetUID,
	})
}
