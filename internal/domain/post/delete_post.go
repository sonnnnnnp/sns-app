package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *PostRepository) DeletePost(ctx context.Context, pID uuid.UUID) error {
	queries := db.New(repo.pool)

	return queries.DeletePost(ctx, pID)
}
