package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *PostRepository) GetPostFavorite(ctx context.Context, uID uuid.UUID, pID uuid.UUID) (*db.GetPostFavoriteRow, error) {
	queries := db.New(repo.pool)

	r, err := queries.GetPostFavorite(ctx, db.GetPostFavoriteParams{
		PostID: pID,
		UserID: uID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrPostFavoriteNotFound
		}
		return nil, err
	}

	return &r, nil
}
