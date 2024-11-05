package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func (repo *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	// u, err := ur.db.User.Get(ctx, id)
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
