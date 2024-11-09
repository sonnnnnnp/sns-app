package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
)

func (uc *UserUsecase) UnfollowUser(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error) {
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

	return &api.SocialConnection{
		Following:  scRow.Following,
		FollowedBy: scRow.FollowedBy,
	}, nil
}
