package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/reverie/internal/http/errors"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *UserUsecase) UnblockUser(ctx context.Context, uID uuid.UUID) error {
	selfUID := ctxhelper.GetUserID(ctx)

	if uID == selfUID {
		return errors.ErrCannotBlockYourself
	}

	return db.New(uc.pool).DeleteUserBlock(ctx, db.DeleteUserBlockParams{
		BlockerID: selfUID,
		BlockedID: uID,
	})
}
