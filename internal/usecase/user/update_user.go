package user_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) UpdateUser(ctx context.Context, id uuid.UUID, body *oapi.UpdateUserJSONBody) (*oapi.User, error) {
	u, err := uu.userRepo.UpdateUser(ctx, id, body)
	if err != nil {
		return nil, err
	}

	return &oapi.User{
		AvatarUrl:   u.AvatarURL,
		Biography:   u.Biography,
		Birthdate:   u.Birthdate,
		CoverUrl:    u.CoverURL,
		CreatedAt:   u.CreatedAt,
		DisplayName: u.DisplayName,
		Id:          u.ID,
		UpdatedAt:   u.UpdatedAt,
	}, nil
}
