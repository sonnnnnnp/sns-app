package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) FollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialConnection, error) {
	uID := ctxhelper.GetUserID(ctx)

	// 自分自身をフォローしようとした場合はエラー
	if uID == targetUID {
		return nil, errors.ErrCannotFollowSelf
	}

	sc, err := uc.userRepo.GetSocialConnection(ctx, uID, targetUID)
	if err != nil {
		return nil, err
	}

	// 既にフォローしている場合はエラー
	if sc.Following {
		return nil, errors.ErrUserAlreadyFollowing
	}

	if err := uc.userRepo.FollowUser(ctx, uID, targetUID); err != nil {
		return nil, err
	}

	return uc.userRepo.GetSocialConnection(ctx, uID, targetUID)
}
