package post

import (
	"context"

	"github.com/google/uuid"
)

func (uc *PostUsecase) DeletePost(ctx context.Context, pID uuid.UUID) error {
	return uc.postRepo.DeletePost(ctx, pID)
}
