package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) FollowUser(ctx context.Context, targetUID uuid.UUID) (*oapi.SocialContext, error) {
	uID := ctxhelper.GetUserID(ctx)

	if err := uu.userRepo.FollowUser(ctx, uID, targetUID); err != nil {
		return nil, err
	}

	u, err := uu.GetUserByID(ctx, targetUID)
	if err != nil {
		return nil, err
	}

	return uu.userRepo.GetSocialContext(ctx, uID, u)
}
