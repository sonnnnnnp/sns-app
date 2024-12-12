package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/reverie/internal/http/adapter/api"
	"github.com/sonnnnnnp/reverie/internal/http/errors"
	"github.com/sonnnnnnp/reverie/internal/http/tools/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *UserUsecase) FollowUser(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error) {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	// 自分自身をフォローしようとした場合はエラー
	if selfUID == uID {
		return nil, errors.ErrCannotFollowYourself
	}

	// ブロックされている場合はエラー
	bs, err := queries.GetBlockStatus(ctx, db.GetBlockStatusParams{
		SelfID:   selfUID,
		TargetID: uID,
	})
	if err != nil {
		return nil, err
	}

	if bs.Blocking || bs.BlockedBy {
		return nil, errors.ErrUserBlockingOrBlockedBy
	}

	sc, err := queries.GetSocialConnection(ctx, db.GetSocialConnectionParams{
		SelfID:   selfUID,
		TargetID: uID,
	})
	if err != nil {
		return nil, err
	}

	// 既にフォローしている場合はエラー
	if sc.Following {
		return nil, errors.ErrUserAlreadyFollowing
	}

	if err := queries.CreateUserFollow(ctx, db.CreateUserFollowParams{
		FollowerID: selfUID,
		FollowedID: uID,
	}); err != nil {
		return nil, err
	}

	scRow, err := queries.GetSocialConnection(ctx, db.GetSocialConnectionParams{
		SelfID:   selfUID,
		TargetID: uID,
	})
	if err != nil {
		return nil, err
	}

	return &api.SocialConnection{
		Following:  scRow.Following,
		FollowedBy: scRow.FollowedBy,
	}, nil
}
