package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/reverie/internal/http/adapter/api"
	internal_errors "github.com/sonnnnnnp/reverie/internal/http/errors"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *UserUsecase) GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]api.UserFollower, error) {
	queries := db.New(uc.pool)

	_, err := queries.GetUserByID(ctx, uID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrUserNotFound
		}
		return nil, err
	}

	// ブロックされている場合はエラー
	bs, err := queries.GetBlockStatus(ctx, db.GetBlockStatusParams{
		SelfID:   ctxhelper.GetUserID(ctx),
		TargetID: uID,
	})
	if err != nil {
		return nil, err
	}

	if bs.Blocking || bs.BlockedBy {
		return nil, internal_errors.ErrUserBlockingOrBlockedBy
	}

	rows, err := queries.GetUserFollowers(ctx, uID)
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
