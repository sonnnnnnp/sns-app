package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) UnfollowUser(ctx context.Context, uID uuid.UUID) (*oapi.SocialConnection, error) {
	selfUID := ctxhelper.GetUserID(ctx)

	if selfUID == uID {
		return nil, errors.ErrCannotFollowYourself
	}

	if err := uc.userRepo.UnfollowUser(ctx, selfUID, uID); err != nil {
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
