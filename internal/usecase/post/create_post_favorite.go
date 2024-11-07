package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
)

func (uc *PostUsecase) CreatePostFavorite(ctx context.Context, pID uuid.UUID) error {
	uID := ctxhelper.GetUserID(ctx)

	_, err := uc.postRepo.GetPostFavorite(ctx, uID, pID)
	if err != nil {
		if errors.Is(err, internal_errors.ErrPostFavoriteNotFound) {
			return uc.postRepo.CreatePostFavorite(ctx, uID, pID)
		}
		return err
	}

	return internal_errors.ErrPostAreadlyFavorited
}
