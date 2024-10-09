package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent/favorite"
)

func (pr *PostRepository) GetPostFavorited(ctx context.Context, uID uuid.UUID, pID uuid.UUID) (bool, error) {
	count, err := pr.db.Favorite.Query().Where(
		favorite.And(
			favorite.UserID(uID),
			favorite.PostID(pID),
		),
	).Count(ctx)
	return count > 0, err
}
