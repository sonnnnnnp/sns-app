package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/reverie/internal/server/http/adapter/api"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *PostUsecase) GetPostFavorites(ctx context.Context, pID uuid.UUID) ([]api.PostFavorite, error) {
	rows, err := db.New(uc.pool).GetPostFavorites(ctx, pID)
	if err != nil {
		return nil, err
	}

	favorites := make([]api.PostFavorite, 0)
	for _, r := range rows {
		favorites = append(favorites, api.PostFavorite{
			CreatedAt: r.PostFavorite.CreatedAt.Time,
			PostId:    r.PostFavorite.PostID,
			User: api.User{
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
