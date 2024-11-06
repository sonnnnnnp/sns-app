package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) FollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialConnection, error) {
	uID := ctxhelper.GetUserID(ctx)

	if uID == targetUID {
		return nil, errors.ErrCannotFollowSelf
	}

	if exists, err := uu.userRepo.GetUserExistence(ctx, targetUID); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.ErrUserNotFound
	}

	if err := uu.userRepo.FollowUser(ctx, uID, targetUID); err != nil {
		return nil, err
	}

	return uu.userRepo.GetSocialConnection(ctx, uID, targetUID)
}
