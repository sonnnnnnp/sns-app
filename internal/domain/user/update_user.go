package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (repo *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, body *oapi.UpdateUserJSONBody) error {
	queries := db.New(repo.pool)

	return queries.UpdateUser(ctx, db.UpdateUserParams{
		UserID:         id,
		Name:           body.Name,
		Nickname:       body.Nickname,
		Biography:      body.Biography,
		AvatarImageUrl: body.AvatarImageUrl,
		BannerImageUrl: body.BannerImageUrl,
		Birthdate:      body.Birthdate,
	})
}
