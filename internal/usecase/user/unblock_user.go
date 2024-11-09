package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *UserUsecase) UnblockUser(ctx context.Context, uID uuid.UUID) error {
	selfUID := ctxhelper.GetUserID(ctx)

	if uID == selfUID {
		return errors.ErrCannotBlockYourself
	}

	return db.New(uc.pool).DeleteUserBlock(ctx, db.DeleteUserBlockParams{
		BlockerID:  selfUID,
		BlockingID: uID,
	})
}
