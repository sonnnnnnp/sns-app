package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) UnfollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialContext, error) {
	uID := ctxhelper.GetUserID(ctx)

	if uID == targetUID {
		return nil, errors.ErrCannotFollowSelf
	}

	if err := uu.userRepo.UnfollowUser(ctx, uID, targetUID); err != nil {
		return nil, err
	}

	return uu.userRepo.GetSocialContext(ctx, uID, targetUID)
}
