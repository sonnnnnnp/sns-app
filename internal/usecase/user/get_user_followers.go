package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]oapi.User, error) {
	// if exists, err := uu.userRepo.GetUserExistence(ctx, uID); err != nil {
	// 	return nil, err
	// } else if !exists {
	// 	return nil, errors.ErrUserNotFound
	// }

	rows, err := uu.userRepo.GetUserFollowers(ctx, uID)
	if err != nil {
		return nil, err
	}

	oapiUsers := make([]oapi.User, len(rows))
	for i, r := range rows {
		oapiUsers[i] = oapi.User{
			AvatarImageUrl: r.AvatarImageUrl,
			Biography:      r.Biography,
			BannerImageUrl: r.BannerImageUrl,
			CreatedAt:      r.CreatedAt.Time,
			Nickname:       r.Nickname,
			Id:             r.ID,
			UpdatedAt:      r.UpdatedAt.Time,
			Name:           r.Name,
		}
	}

	return oapiUsers, nil
}
