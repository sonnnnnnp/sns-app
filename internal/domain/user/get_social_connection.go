package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (repo *UserRepository) GetSocialConnection(ctx context.Context, selfID uuid.UUID, targetUID uuid.UUID) (*oapi.SocialConnection, error) {
	queries := db.New(repo.pool)

	row, err := queries.GetSocialConnection(ctx, db.GetSocialConnectionParams{
		SelfID:   selfID,
		TargetID: targetUID,
	})
	if err != nil {
		return nil, err
	}

	return &oapi.SocialConnection{
		Following:  row.Following,
		FollowedBy: row.FollowedBy,
	}, nil
}
