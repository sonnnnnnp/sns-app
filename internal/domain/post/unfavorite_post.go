package post

import (
	"context"

	"github.com/google/uuid"
)

func (repo *PostRepository) UnfavoritePost(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error {
	// _, err := pr.db.Favorite.Delete().Where(
	// 	favorite.And(
	// 		favorite.UserID(uID),
	// 		favorite.PostID(pID),
	// 	),
	// ).Exec(ctx)

	// return err

	return nil
}
