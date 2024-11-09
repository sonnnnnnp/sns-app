package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
)

func (uc *UserUsecase) UnblockUser(ctx context.Context, uID uuid.UUID) error {
	selfUID := ctxhelper.GetUserID(ctx)

	if uID == selfUID {
		return errors.ErrCannotBlockSelf
	}

	return uc.userRepo.UnblockUser(ctx, selfUID, uID)
}
