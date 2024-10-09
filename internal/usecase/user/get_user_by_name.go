package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) GetUserByName(ctx context.Context, name string) (*oapi.User, error) {
	uID := ctxhelper.GetUserID(ctx)

	u, err := uu.userRepo.GetUserByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var sc *oapi.SocialContext
	if u.ID != uID {
		sc, err = uu.userRepo.GetSocialContext(ctx, uID, u.ID)
		if err != nil {
			return nil, err
		}
	}

	return &oapi.User{
		Id:             u.ID,
		Name:           u.Name,
		Nickname:       u.Nickname,
		AvatarImageUrl: &u.AvatarImageURL,
		BannerImageUrl: &u.BannerImageURL,
		Biography:      &u.Biography,
		SocialContext:  sc,
		UpdatedAt:      u.UpdatedAt,
		CreatedAt:      u.CreatedAt,
	}, nil
}
