package user

import (
	"context"

	"github.com/sonnnnnnp/reverie/internal/adapter/api"
	"github.com/sonnnnnnp/reverie/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *UserUsecase) GetUserBlocking(ctx context.Context) ([]api.User, error) {
	rows, err := db.New(uc.pool).GetUserBlocking(ctx, ctxhelper.GetUserID(ctx))
	if err != nil {
		return nil, err
	}

	blocking := make([]api.User, 0)
	for _, u := range rows {
		blocking = append(blocking, api.User{
			AvatarImageUrl: u.AvatarImageUrl,
			BannerImageUrl: u.BannerImageUrl,
			Biography:      u.Biography,
			CreatedAt:      u.CreatedAt.Time,
			Nickname:       u.Nickname,
			Id:             u.ID,
			UpdatedAt:      u.UpdatedAt.Time,
			Name:           u.Name,
		})
	}

	return blocking, nil
}
