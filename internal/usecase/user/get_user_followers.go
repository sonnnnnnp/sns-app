package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
)

func (uc *UserUsecase) GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]api.UserFollower, error) {
	u, err := uc.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.ErrUserNotFound
	}

	// ブロックされている場合はエラー
	bs, err := uc.userRepo.GetBlockStatus(ctx, ctxhelper.GetUserID(ctx), uID)
	if err != nil {
		return nil, err
	}

	if bs.Blocking || bs.BlockedBy {
		return nil, errors.ErrUserBlockingOrBlockedBy
	}

	rows, err := uc.userRepo.GetUserFollowers(ctx, uID)
	if err != nil {
		return nil, err
	}

	followers := make([]api.UserFollower, 0)
	for _, r := range rows {
		followers = append(followers, api.UserFollower{
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
