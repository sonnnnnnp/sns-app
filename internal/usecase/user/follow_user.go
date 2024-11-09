package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) FollowUser(ctx context.Context, uID uuid.UUID) (*oapi.SocialConnection, error) {
	selfUID := ctxhelper.GetUserID(ctx)

	// 自分自身をフォローしようとした場合はエラー
	if selfUID == uID {
		return nil, errors.ErrCannotFollowYourself
	}

	// ブロックされている場合はエラー
	bs, err := uc.userRepo.GetBlockStatus(ctx, selfUID, uID)
	if err != nil {
		return nil, err
	}

	if bs.Blocking || bs.BlockedBy {
		return nil, errors.ErrUserBlockingOrBlockedBy
	}

	sc, err := uc.userRepo.GetSocialConnection(ctx, selfUID, uID)
	if err != nil {
		return nil, err
	}

	// 既にフォローしている場合はエラー
	if sc.Following {
		return nil, errors.ErrUserAlreadyFollowing
	}

	if err := uc.userRepo.FollowUser(ctx, selfUID, uID); err != nil {
		return nil, err
	}

	scRow, err := uc.userRepo.GetSocialConnection(ctx, selfUID, uID)
	if err != nil {
		return nil, err
	}

	return &oapi.SocialConnection{
		Following:  scRow.Following,
		FollowedBy: scRow.FollowedBy,
	}, nil
}
