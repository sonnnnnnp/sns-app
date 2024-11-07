package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (*oapi.User, error) {
	uID := ctxhelper.GetUserID(ctx)

	u, err := uc.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.ErrUserNotFound
	}

	var sc *oapi.SocialConnection
	if u.ID != uID {
		sc, err = uc.userRepo.GetSocialConnection(ctx, uID, u.ID)
		if err != nil {
			return nil, err
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
		UpdatedAt:        u.UpdatedAt.Time,
		CreatedAt:        u.CreatedAt.Time,
	}, nil
}
