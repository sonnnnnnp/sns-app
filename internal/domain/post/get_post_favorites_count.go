package post

import (
	"context"

	"github.com/google/uuid"
)

func (repo *PostRepository) GetPostFavoritesCount(ctx context.Context, pID uuid.UUID) (int, error) {
	// return pr.db.Favorite.Query().Where(favorite.PostID(pID)).Count(ctx)

	return 0, nil
}
