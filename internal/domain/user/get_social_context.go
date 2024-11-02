package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (ur *UserRepository) GetSocialContext(ctx context.Context, selfID uuid.UUID, targetUID uuid.UUID) (*oapi.SocialContext, error) {
	// isFollowedBy, err := ur.db.User.Query().Where(user.ID(targetUID)).QueryFollowing().Exist(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// isFollowing, err := ur.db.User.Query().Where(user.ID(targetUID)).QueryFollowers().Exist(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// return &oapi.SocialContext{
	// 	FollowedBy: isFollowedBy,
	// 	Following:  isFollowing,
	// }, nil

	return nil, nil
}
