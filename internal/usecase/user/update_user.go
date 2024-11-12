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

	row, err := queries.GetUserByID(ctx, selfUID)
	if err != nil {
		return nil, err
	}

	return &api.User{
		Id:             row.User.ID,
		Name:           row.User.Name,
		Nickname:       row.User.Nickname,
		AvatarImageUrl: row.User.AvatarImageUrl,
		BannerImageUrl: row.User.BannerImageUrl,
		Biography:      row.User.Biography,
		UpdatedAt:      row.User.UpdatedAt.Time,
		CreatedAt:      row.User.CreatedAt.Time,
	}, nil
}
