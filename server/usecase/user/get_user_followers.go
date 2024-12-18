package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/infra/db"
	"github.com/sonnnnnnp/reverie/server/pkg/ctxhelper"
	internal_errors "github.com/sonnnnnnp/reverie/server/pkg/errors"
)

func (uc *UserUsecase) GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]api.UserFollower, error) {
	queries := db.New(uc.pool)

	// TODO: これいる？
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

	followers := make([]api.UserFollower, 0)
	for _, r := range rows {
		followers = append(followers, api.UserFollower{
			AvatarImageUrl: r.AvatarImageUrl,
			Biography:      r.Biography,
			BannerImageUrl: r.BannerImageUrl,
			CreatedAt:      &r.CreatedAt.Time,
			FollowedAt:     r.FollowedAt.Time,
			Nickname:       r.Nickname,
			Id:             r.ID,
			CustomId:           r.CustomID,
		})
	}

	return followers, nil
}
