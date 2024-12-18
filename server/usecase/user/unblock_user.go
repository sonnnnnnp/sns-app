package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/reverie/server/infra/db"
	"github.com/sonnnnnnp/reverie/server/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/server/pkg/errors"
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
