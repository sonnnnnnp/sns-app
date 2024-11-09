package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
)

func (uc *UserUsecase) GetUserByID(ctx context.Context, uID uuid.UUID) (*api.User, error) {
	selfUID := ctxhelper.GetUserID(ctx)

	u, err := uc.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.ErrUserNotFound
	}

	var sc *api.SocialConnection
	var bs *api.BlockStatus
	if u.ID != selfUID {
		scRow, err := uc.userRepo.GetSocialConnection(ctx, selfUID, u.ID)
		if err != nil {
			return nil, err
		}
		sc = &api.SocialConnection{
			Following:  scRow.Following,
			FollowedBy: scRow.FollowedBy,
		}

		bsRow, err := uc.userRepo.GetBlockStatus(ctx, selfUID, u.ID)
		if err != nil {
			return nil, err
		}
		bs = &api.BlockStatus{
			BlockedBy: bsRow.BlockedBy,
			Blocking:  bsRow.Blocking,
		}
	}

	return &api.User{
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
