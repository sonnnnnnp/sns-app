package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (*oapi.User, error) {
	uID := ctxhelper.GetUserID(ctx)

	u, err := uu.userRepo.GetUserByID(ctx, id)
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
