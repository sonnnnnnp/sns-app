package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) GetUserBlocking(ctx context.Context, uID uuid.UUID) ([]db.GetUserBlockingRow, error) {
	return db.New(repo.pool).GetUserBlocking(ctx, uID)
}
