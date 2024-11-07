package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *PostRepository) GetPostFavorites(ctx context.Context, pID uuid.UUID) ([]db.GetPostFavoritesRow, error) {
	queries := db.New(repo.pool)

	rows, err := queries.GetPostFavorites(ctx, pID)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
