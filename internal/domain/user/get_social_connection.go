package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetSocialConnection(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) (db.GetSocialConnectionRow, error) {
	queries := db.New(repo.pool)

	return queries.GetSocialConnection(ctx, db.GetSocialConnectionParams{
		SelfID:   selfUID,
		TargetID: targetUID,
	})
}
