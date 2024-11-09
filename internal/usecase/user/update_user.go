package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) UpdateUser(ctx context.Context, body *oapi.UpdateUserJSONBody) (*oapi.User, error) {
	selfUID := ctxhelper.GetUserID(ctx)

	if err := uc.userRepo.UpdateUser(ctx, selfUID, body); err != nil {
		return nil, err
	}

	u, err := uc.userRepo.GetUserByID(ctx, selfUID)
	if err != nil {
		return nil, err
	}

	return &oapi.User{
		AvatarImageUrl: u.AvatarImageUrl,
		BannerImageUrl: u.BannerImageUrl,
		Biography:      u.Biography,
		CreatedAt:      u.CreatedAt.Time,
		Name:           u.Name,
		Nickname:       u.Nickname,
		Id:             u.ID,
		UpdatedAt:      u.UpdatedAt.Time,
	}, nil
}
