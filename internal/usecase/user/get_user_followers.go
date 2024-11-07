package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]oapi.UserFollower, error) {
	u, err := uc.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.ErrUserNotFound
	}

	rows, err := uc.userRepo.GetUserFollowers(ctx, uID)
	if err != nil {
		return nil, err
	}

	followers := make([]oapi.UserFollower, 0)
	for _, r := range rows {
		followers = append(followers, oapi.UserFollower{
			AvatarImageUrl: r.AvatarImageUrl,
			Biography:      r.Biography,
			BannerImageUrl: r.BannerImageUrl,
			CreatedAt:      r.CreatedAt.Time,
			FollowedAt:     r.FollowedAt.Time,
			Nickname:       r.Nickname,
			Id:             r.ID,
			UpdatedAt:      r.UpdatedAt.Time,
			Name:           r.Name,
		})
	}

	return followers, nil
}
