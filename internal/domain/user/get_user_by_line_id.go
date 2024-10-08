package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/ent/user"
)

func (ur *UserRepository) GetUserByLineID(ctx context.Context, lineID string) (*ent.User, error) {
	u, err := ur.db.User.Query().Where(user.LineID(lineID)).First(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, errors.ErrUserNotFound
		default:
			return nil, err
		}
	}

	return u, nil
}
