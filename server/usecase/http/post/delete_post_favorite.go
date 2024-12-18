package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/reverie/server/infra/db"
	"github.com/sonnnnnnp/reverie/server/pkg/ctxhelper"
)

func (uc *PostUsecase) DeletePostFavorite(ctx context.Context, pID uuid.UUID) error {
	return db.New(uc.pool).DeletePostFavorite(ctx, db.DeletePostFavoriteParams{
		UserID: ctxhelper.GetUserID(ctx),
		PostID: pID,
	})
}
