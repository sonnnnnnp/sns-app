package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (ur *UserRepository) GetUserFollowers(ctx context.Context, u *ent.User) ([]*ent.User, error) {
	followers := u.QueryFollowers().AllX(ctx)
	return followers, nil
}
