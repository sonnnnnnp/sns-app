package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *PostUsecase) DeletePostFavorite(ctx context.Context, pID uuid.UUID) error {
	return db.New(uc.pool).DeletePostFavorite(ctx, db.DeletePostFavoriteParams{
		UserID: ctxhelper.GetUserID(ctx),
		PostID: pID,
	})
}
