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

	if body.Name != nil {
		query.SetName(*body.Name)
	}
	if body.Nickname != nil {
		query.SetNickname(*body.Nickname)
	}
	if body.Biography != nil {
		query.SetBiography(*body.Biography)
	}
	if body.AvatarImageUrl != nil {
		query.SetAvatarImageURL(*body.AvatarImageUrl)
	}
	if body.BannerImageUrl != nil {
		query.SetBannerImageURL(*body.BannerImageUrl)
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
