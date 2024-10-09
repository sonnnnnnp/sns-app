package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/ent/user"
)

func (ur *UserRepository) GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]*ent.User, error) {
	return ur.db.User.Query().Where(user.ID(uID)).QueryFollowing().All(ctx)
}
