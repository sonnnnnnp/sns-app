package user

import (
	"context"

	"github.com/google/uuid"
)

func (repo *UserRepository) GetUserExistence(ctx context.Context, uID uuid.UUID) (bool, error) {
	// return ur.db.User.Query().Where(user.ID(uID)).Exist(ctx)

	return false, nil
}
