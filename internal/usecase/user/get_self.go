package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *UserUsecase) GetSelf(ctx context.Context) (*api.User, error) {
	selfUID := ctxhelper.GetUserID(ctx)

	queries := db.New(uc.pool)

	row, err := queries.GetUserByID(ctx, selfUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrUserNotFound
		}
		return nil, err
	}

	scRow, err := queries.GetSocialConnection(ctx, db.GetSocialConnectionParams{
		SelfID:   selfUID,
		TargetID: row.User.ID,
	})
	if err != nil {
		return nil, err
	}

	bsRow, err := queries.GetBlockStatus(ctx, db.GetBlockStatusParams{
		SelfID:   selfUID,
		TargetID: row.User.ID,
	})
	if err != nil {
		return nil, err
	}

	return &api.User{
		Id:             row.User.ID,
		Name:           row.User.Name,
		Nickname:       row.User.Nickname,
		AvatarImageUrl: row.User.AvatarImageUrl,
		BannerImageUrl: row.User.BannerImageUrl,
		Biography:      row.User.Biography,
		BlockStatus: &api.BlockStatus{
			BlockedBy: bsRow.BlockedBy,
			Blocking:  bsRow.Blocking,
		},
		SocialConnection: &api.SocialConnection{
			Following:  scRow.Following,
			FollowedBy: scRow.FollowedBy,
		},
		SocialEngagement: &api.SocialEngagement{
			FollowingCount: int(row.FollowingCount),
			FollowersCount: int(row.FollowersCount),
			PostsCount:     int(row.PostsCount),
			MediaCount:     int(row.MediaCount),
			FavoritesCount: int(row.FavoritesCount),
		},
		UpdatedAt: row.User.UpdatedAt.Time,
		CreatedAt: row.User.CreatedAt.Time,
	}, nil
}
