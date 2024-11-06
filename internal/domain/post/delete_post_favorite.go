package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *PostRepository) DeletePostFavorite(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error {
	queries := db.New(repo.pool)

	if err := queries.DeletePostFavorite(ctx, db.DeletePostFavoriteParams{
		UserID: uID,
		PostID: pID,
	}); err != nil {
		return err
	}

	return nil
}
