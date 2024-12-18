package user

import (
	"context"

	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/infra/db"
	"github.com/sonnnnnnp/reverie/server/pkg/ctxhelper"
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
		Id:             u.ID,
		Name:           u.Name,
		Nickname:       u.Nickname,
		AvatarImageUrl: u.AvatarImageUrl,
		BannerImageUrl: u.BannerImageUrl,
		Biography:      u.Biography,
		UpdatedAt:      u.UpdatedAt.Time,
		CreatedAt:      u.CreatedAt.Time,
	}, nil
}
