package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent/user"
)

func (ur *UserRepository) GetUserExistence(ctx context.Context, uID uuid.UUID) (bool, error) {
	return ur.db.User.Query().Where(user.ID(uID)).Exist(ctx)
}
