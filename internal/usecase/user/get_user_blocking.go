package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) GetUserBlocking(ctx context.Context) ([]oapi.User, error) {
	rows, err := uc.userRepo.GetUserBlocking(ctx, ctxhelper.GetUserID(ctx))
	if err != nil {
		return nil, err
	}

	blocking := make([]oapi.User, 0)
	for _, u := range rows {
		blocking = append(blocking, oapi.User{
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
