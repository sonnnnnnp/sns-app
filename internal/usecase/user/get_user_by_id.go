package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *UserUsecase) GetUserByID(ctx context.Context, uID uuid.UUID) (*api.User, error) {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	row, err := queries.GetUserByID(ctx, uID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrUserNotFound
		}
		return nil, err
	}

	var scRow db.GetSocialConnectionRow
	var bsRow db.GetBlockStatusRow
	if row.User.ID != selfUID {
		scRow, err = queries.GetSocialConnection(ctx, db.GetSocialConnectionParams{
			SelfID:   selfUID,
			TargetID: row.User.ID,
		})
		if err != nil {
			return nil, err
		}

		bsRow, err = queries.GetBlockStatus(ctx, db.GetBlockStatusParams{
			SelfID:   selfUID,
			TargetID: row.User.ID,
		})
		if err != nil {
			return nil, err
		}
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
