package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *UserUsecase) RemoveUserFromFollowers(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error) {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	if selfUID == uID {
		return nil, errors.ErrCannotFollowYourself
	}

	if err := queries.DeleteUserFollow(ctx, db.DeleteUserFollowParams{
		FollowerID:  uID,
		FollowingID: selfUID,
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
