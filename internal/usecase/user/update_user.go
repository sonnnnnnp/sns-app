package user

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
		AvatarImageUrl: &u.AvatarImageURL,
		BannerImageUrl: &u.BannerImageURL,
		Biography:      &u.Biography,
		CreatedAt:      u.CreatedAt,
		Nickname:       u.Nickname,
		Id:             u.ID,
		UpdatedAt:      u.UpdatedAt,
	}, nil
}
