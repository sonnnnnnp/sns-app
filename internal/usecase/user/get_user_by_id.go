package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) GetUserByID(ctx context.Context, uID uuid.UUID) (*oapi.User, error) {
	selfUID := ctxhelper.GetUserID(ctx)

	u, err := uc.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.ErrUserNotFound
	}

	var sc *oapi.SocialConnection
	var bs *oapi.BlockStatus
	if u.ID != selfUID {
		scRow, err := uc.userRepo.GetSocialConnection(ctx, selfUID, u.ID)
		if err != nil {
			return nil, err
		}
		sc = &oapi.SocialConnection{
			Following:  scRow.Following,
			FollowedBy: scRow.FollowedBy,
		}

		bsRow, err := uc.userRepo.GetBlockStatus(ctx, selfUID, u.ID)
		if err != nil {
			return nil, err
		}
		bs = &oapi.BlockStatus{
			BlockedBy: bsRow.BlockedBy,
			Blocking:  bsRow.Blocking,
		}
	}

	return &oapi.User{
		Id:               u.ID,
		Name:             u.Name,
		Nickname:         u.Nickname,
		AvatarImageUrl:   u.AvatarImageUrl,
		BannerImageUrl:   u.BannerImageUrl,
		Biography:        u.Biography,
		SocialConnection: sc,
		BlockStatus:      bs,
		UpdatedAt:        u.UpdatedAt.Time,
		CreatedAt:        u.CreatedAt.Time,
	}, nil
}
