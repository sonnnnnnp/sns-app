package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/ent/user"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (ur *UserRepository) GetSocialContext(ctx context.Context, selfID uuid.UUID, target *ent.User) (*oapi.SocialContext, error) {
	isFollowedBy, err := target.QueryFollowers().Where(user.IDEQ(selfID)).Exist(ctx)
	if err != nil {
		return nil, err
	}

	isFollowing, err := target.QueryFollowing().Where(user.IDEQ(selfID)).Exist(ctx)
	if err != nil {
		return nil, err
	}

	return &oapi.SocialContext{
		FollowedBy: isFollowedBy,
		Following:  isFollowing,
	}, nil
}
