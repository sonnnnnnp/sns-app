package timeline

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *TimelineUsecase) GetTimeline(ctx context.Context, params *oapi.GetTimelineParams) (posts []oapi.Post, nextCursor uuid.UUID, err error) {
	// if parmas.UserId is not nil, check if not blocked by the target user

	rows, nextCursor, err := uc.postRepo.GetTimeline(ctx, params.Limit, params.Cursor, params.UserId)
	if err != nil {
		return nil, uuid.Nil, err
	}

	posts = make([]oapi.Post, 0)
	for _, r := range rows {
		posts = append(posts, oapi.Post{
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
			CreatedAt:      r.Post.CreatedAt.Time,
			Favorited:      r.Favorited,
			FavoritesCount: int(r.FavoritesCount),
			Id:             r.Post.ID,
			Text:           r.Post.Text,
			UpdatedAt:      r.Post.UpdatedAt.Time,
		})
	}

	return posts, nextCursor, nil
}
