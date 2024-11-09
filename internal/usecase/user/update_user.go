package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *UserUsecase) UpdateUser(ctx context.Context, body *api.UpdateUserJSONBody) (*api.User, error) {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	if err := queries.UpdateUser(ctx, db.UpdateUserParams{
		UserID:         selfUID,
		Name:           body.Name,
		Nickname:       body.Nickname,
		Biography:      body.Biography,
		AvatarImageUrl: body.AvatarImageUrl,
		BannerImageUrl: body.BannerImageUrl,
		Birthdate:      body.Birthdate,
	}); err != nil {
		return nil, err
	}

	u, err := queries.GetUserByID(ctx, selfUID)
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
