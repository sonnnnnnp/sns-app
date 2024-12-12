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
