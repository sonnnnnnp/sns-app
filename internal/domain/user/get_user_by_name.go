package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (repo *UserRepository) GetUserByName(ctx context.Context, name string) (*ent.User, error) {
	// u, err := ur.db.User.Query().Where(user.Name(name)).First(ctx)
	// if err != nil {
	// 	switch {
	// 	case ent.IsNotFound(err):
	// 		return nil, errors.ErrUserNotFound
	// 	default:
	// 		return nil, err
	// 	}
	// }

	// return u, nil

	return nil, nil
}
