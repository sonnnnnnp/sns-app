package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *UserUsecase) BlockUser(ctx context.Context, uID uuid.UUID) error {
	// TODO: どちらかがフォロー中の場合はフォロー解除する
	// トランザクションを利用する

	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	// 自分自身をブロックしようとした場合はエラー
	if uID == selfUID {
		return errors.ErrCannotBlockYourself
	}

	bs, err := queries.GetBlockStatus(ctx, db.GetBlockStatusParams{
		SelfID:   selfUID,
		TargetID: uID,
	})
	if err != nil {
		return err
	}

	// 既にブロックしている場合はエラー
	if bs.Blocking {
		return errors.ErrUserAlreadyBlocking
	}

	return queries.CreateUserBlock(ctx, db.CreateUserBlockParams{
		BlockerID: selfUID,
		BlockedID: uID,
	})
}
