package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/reverie/server/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/server/infra/db"
	internal_errors "github.com/sonnnnnnp/reverie/server/pkg/errors"
)

func (uc *PostUsecase) CreatePostFavorite(ctx context.Context, pID uuid.UUID) error {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	_, err := queries.GetPostFavorite(ctx, db.GetPostFavoriteParams{
		PostID: pID,
		UserID: selfUID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return queries.CreatePostFavorite(ctx, db.CreatePostFavoriteParams{
				UserID: selfUID,
				PostID: pID,
			})
		}
		return err
	}

	return internal_errors.ErrPostAreadlyFavorited
}
