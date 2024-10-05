package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (ur *UserRepository) GetUserFollowing(ctx context.Context, u *ent.User) ([]*ent.User, error) {
	following := u.QueryFollowing().AllX(ctx)
	return following, nil
}
