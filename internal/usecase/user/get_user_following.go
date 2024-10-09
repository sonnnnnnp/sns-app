package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]oapi.User, error) {
	if exists, err := uu.userRepo.GetUserExistence(ctx, uID); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.ErrUserNotFound
	}

	users, err := uu.userRepo.GetUserFollowing(ctx, uID)
	if err != nil {
		return nil, err
	}

	oapiUsers := make([]oapi.User, len(users))
	for i, u := range users {
		oapiUsers[i] = oapi.User{
			AvatarImageUrl: &u.AvatarImageURL,
			Biography:      &u.Biography,
			BannerImageUrl: &u.BannerImageURL,
			CreatedAt:      u.CreatedAt,
			Nickname:       u.Nickname,
			Id:             u.ID,
			UpdatedAt:      u.UpdatedAt,
			Name:           u.Name,
		}
	}

	return oapiUsers, nil
}
