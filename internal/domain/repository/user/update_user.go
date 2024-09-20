package user_repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (ur *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, body *oapi.UpdateUserJSONBody) (*ent.User, error) {
	builder := ur.db.User.UpdateOneID(id)

	if body.Username != nil {
		builder.SetUsername(*body.Username)
	}
	if body.DisplayName != nil {
		builder.SetDisplayName(*body.DisplayName)
	}
	if body.Biography != nil {
		builder.SetBiography(*body.Biography)
	}
	if body.AvatarUrl != nil {
		builder.SetAvatarURL(*body.AvatarUrl)
	}
	if body.CoverUrl != nil {
		builder.SetCoverURL(*body.CoverUrl)
	}
	if body.Birthdate != nil {
		builder.SetBirthdate(*body.Birthdate)
	}

	builder.SetUpdatedAt(time.Now())

	u, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
