package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *PostUsecase) GetPostFavorites(ctx context.Context, pID uuid.UUID) (favorites []oapi.PostFavorite, err error) {
	rows, err := uc.postRepo.GetPostFavorites(ctx, pID)
	if err != nil {
		return nil, err
	}

	for _, r := range rows {
		favorites = append(favorites, oapi.PostFavorite{
			CreatedAt: r.PostFavorite.CreatedAt.Time,
			PostId:    r.PostFavorite.PostID,
			User: oapi.User{
				Id:             r.User.ID,
				Name:           r.User.Name,
				Nickname:       r.User.Nickname,
				AvatarImageUrl: r.User.AvatarImageUrl,
				BannerImageUrl: r.User.BannerImageUrl,
				Biography:      r.User.Biography,
				CreatedAt:      r.User.CreatedAt.Time,
				UpdatedAt:      r.User.UpdatedAt.Time,
			},
		})
	}

	return favorites, nil
}
