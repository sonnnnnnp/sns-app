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

	u, err := queries.GetUserByID(ctx, uID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrUserNotFound
		}
		return nil, err
	}

	var sc *api.SocialConnection
	var bs *api.BlockStatus
	if u.ID != selfUID {
		scRow, err := queries.GetSocialConnection(ctx, db.GetSocialConnectionParams{
			SelfID:   selfUID,
			TargetID: u.ID,
		})
		if err != nil {
			return nil, err
		}
		sc = &api.SocialConnection{
			Following:  scRow.Following,
			FollowedBy: scRow.FollowedBy,
		}

		bsRow, err := queries.GetBlockStatus(ctx, db.GetBlockStatusParams{
			SelfID:   selfUID,
			TargetID: u.ID,
		})
		if err != nil {
			return nil, err
		}
		bs = &api.BlockStatus{
			BlockedBy: bsRow.BlockedBy,
			Blocking:  bsRow.Blocking,
		}
	}

	return &api.User{
		Id:               u.ID,
		Name:             u.Name,
		Nickname:         u.Nickname,
		AvatarImageUrl:   u.AvatarImageUrl,
		BannerImageUrl:   u.BannerImageUrl,
		Biography:        u.Biography,
		SocialConnection: sc,
		BlockStatus:      bs,
		UpdatedAt:        u.UpdatedAt.Time,
		CreatedAt:        u.CreatedAt.Time,
	}, nil
}
