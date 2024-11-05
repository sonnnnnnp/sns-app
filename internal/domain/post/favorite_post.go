package post

import (
	"context"

	"github.com/google/uuid"
)

func (repo *PostRepository) FavoritePost(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error {
	// _, err := pr.db.Favorite.Create().SetUserID(uID).SetPostID(pID).Save(ctx)
	// if err != nil {
	// 	if ent.IsConstraintError(err) {
	// 		return errors.ErrPostAreadlyFavorited
	// 	}
	// 	return err
	// }

	return nil
}
