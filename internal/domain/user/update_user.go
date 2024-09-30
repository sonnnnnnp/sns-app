package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (ur *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, body *oapi.UpdateUserJSONBody) (*ent.User, error) {
	query := ur.db.User.UpdateOneID(id)

	if body.Username != nil {
		query.SetUsername(*body.Username)
	}
	if body.DisplayName != nil {
		query.SetDisplayName(*body.DisplayName)
	}
	if body.Biography != nil {
		query.SetBiography(*body.Biography)
	}
	if body.AvatarUrl != nil {
		query.SetAvatarURL(*body.AvatarUrl)
	}
	if body.CoverUrl != nil {
		query.SetCoverURL(*body.CoverUrl)
	}
	if body.Birthdate != nil {
		query.SetBirthdate(*body.Birthdate)
	}

	query.SetUpdatedAt(time.Now())

	u, err := query.Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
