package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, body *api.UpdateUserJSONBody) error {
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
