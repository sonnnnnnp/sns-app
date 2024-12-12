package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/reverie/internal/http/adapter/api"
	"github.com/sonnnnnnp/reverie/internal/http/errors"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *UserUsecase) RemoveUserFromFollowers(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error) {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	if selfUID == uID {
		return nil, errors.ErrCannotFollowYourself
	}

	if err := queries.DeleteUserFollow(ctx, db.DeleteUserFollowParams{
		FollowerID: uID,
		FollowedID: selfUID,
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
