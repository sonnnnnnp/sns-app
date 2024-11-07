package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *PostUsecase) GetPostByID(ctx context.Context, id uuid.UUID) (*oapi.Post, error) {
	r, err := uc.postRepo.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if r == nil {
		return nil, errors.ErrPostNotFound
	}

	return &oapi.Post{
		Author: oapi.User{
			AvatarImageUrl: r.User.AvatarImageUrl,
			BannerImageUrl: r.User.BannerImageUrl,
			Biography:      r.User.Biography,
			CreatedAt:      r.User.CreatedAt.Time,
			Id:             r.User.ID,
			Name:           r.User.Name,
			Nickname:       r.User.Nickname,
			UpdatedAt:      r.User.UpdatedAt.Time,
		},
		Id:             r.Post.ID,
		Text:           r.Post.Text,
		Favorited:      r.Favorited,
		FavoritesCount: int(r.FavoritesCount),
		CreatedAt:      r.Post.CreatedAt.Time,
		UpdatedAt:      r.Post.UpdatedAt.Time,
	}, nil
}
