package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/ent/user"
)

func (ur *UserRepository) GetUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	u, err := ur.db.User.Query().Where(user.Username(username)).First(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, nil
		default:
			return nil, err
		}
	}

	return u, nil
}
