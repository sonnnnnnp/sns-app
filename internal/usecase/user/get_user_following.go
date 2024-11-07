package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]oapi.UserFollower, error) {
	u, err := uc.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.ErrUserNotFound
	}

	rows, err := uc.userRepo.GetUserFollowing(ctx, uID)
	if err != nil {
		return nil, err
	}

	following := make([]oapi.UserFollower, 0)
	for _, u := range rows {
		following = append(following, oapi.UserFollower{
			AvatarImageUrl: u.AvatarImageUrl,
			BannerImageUrl: u.BannerImageUrl,
			Biography:      u.Biography,
			CreatedAt:      u.CreatedAt.Time,
			FollowedAt:     u.FollowedAt.Time,
			Nickname:       u.Nickname,
			Id:             u.ID,
			UpdatedAt:      u.UpdatedAt.Time,
			Name:           u.Name,
		})
	}

	return following, nil
}
