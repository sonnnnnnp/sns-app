package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (pr *PostRepository) FavoritePost(ctx context.Context, uID uuid.UUID, pID uuid.UUID) error {
	_, err := pr.db.Favorite.Create().SetUserID(uID).SetPostID(pID).Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return errors.ErrPostAreadlyFavorited
		}
		return err
	}

	return nil
}
