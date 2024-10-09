package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent/favorite"
)

func (pr *PostRepository) GetPostFavoritesCount(ctx context.Context, pID uuid.UUID) (int, error) {
	return pr.db.Favorite.Query().Where(favorite.PostID(pID)).Count(ctx)
}
