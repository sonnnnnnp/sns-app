package user_repository

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent/user"
)

func (ur *UserRepository) GetUserByLineID(ctx context.Context, lineID string) (*ent.User, error) {
	u, err := ur.db.User.Query().Where(user.LineID(lineID)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return u, nil
}
