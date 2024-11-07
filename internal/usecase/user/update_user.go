package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *UserUsecase) UpdateUser(ctx context.Context, id uuid.UUID, body *oapi.UpdateUserJSONBody) (*oapi.User, error) {
	if err := uc.userRepo.UpdateUser(ctx, id, body); err != nil {
		return nil, err
	}

	u, err := uc.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.ErrUserNotFound
	}

	return &oapi.User{
		AvatarImageUrl: u.AvatarImageUrl,
		BannerImageUrl: u.BannerImageUrl,
		Biography:      u.Biography,
		CreatedAt:      u.CreatedAt.Time,
		Name:           u.Name,
		Nickname:       u.Nickname,
		Id:             u.ID,
		UpdatedAt:      u.UpdatedAt.Time,
	}, nil
}
