package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
)

func (uc *UserUsecase) BlockUser(ctx context.Context, uID uuid.UUID) error {
	selfUID := ctxhelper.GetUserID(ctx)

	// 自分自身をブロックしようとした場合はエラー
	if uID == selfUID {
		return errors.ErrCannotBlockYourself
	}

	bs, err := uc.userRepo.GetBlockStatus(ctx, selfUID, uID)
	if err != nil {
		return err
	}

	// 既にブロックしている場合はエラー
	if bs.Blocking {
		return errors.ErrUserAlreadyBlocking
	}

	return uc.userRepo.BlockUser(ctx, selfUID, uID)
}
