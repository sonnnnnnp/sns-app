package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
)

func (uc *UserUsecase) UpdateUser(ctx context.Context, body *api.UpdateUserJSONBody) (*api.User, error) {
	selfUID := ctxhelper.GetUserID(ctx)

	if err := uc.userRepo.UpdateUser(ctx, selfUID, body); err != nil {
		return nil, err
	}

	u, err := uc.userRepo.GetUserByID(ctx, selfUID)
	if err != nil {
		return nil, err
	}

	return &api.User{
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
