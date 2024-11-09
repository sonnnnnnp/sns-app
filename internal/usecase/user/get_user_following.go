package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
)

func (uc *UserUsecase) GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]api.UserFollower, error) {
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

	rows, err := uc.userRepo.GetUserFollowing(ctx, uID)
	if err != nil {
		return nil, err
	}

	following := make([]api.UserFollower, 0)
	for _, u := range rows {
		following = append(following, api.UserFollower{
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
