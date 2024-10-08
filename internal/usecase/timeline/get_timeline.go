package timeline

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (tu *TimelineUsecase) GetTimeline(ctx context.Context, params *oapi.GetTimelineParams) (posts []oapi.Post, nextCursor uuid.UUID, err error) {
	uID := ctxhelper.GetUserID(ctx)

	// if parmas.UserId is not nil, check if not blocked by the target user

	entPosts, nextCursor, err := tu.postRepo.GetPosts(ctx, params.Limit, params.Cursor, params.UserId)
	if err != nil {
		return nil, uuid.Nil, err
	}

	for _, p := range entPosts {
		favoritesCount, err := tu.postRepo.GetPostFavoritesCount(ctx, p.ID)
		if err != nil {
			return nil, uuid.Nil, err
		}

		favorited, err := tu.postRepo.GetPostFavorited(ctx, uID, p.ID)
		if err != nil {
			return nil, uuid.Nil, err
		}

		posts = append(posts, oapi.Post{
			Author: oapi.User{
				AvatarImageUrl: &p.Edges.Author.AvatarImageURL,
				BannerImageUrl: &p.Edges.Author.BannerImageURL,
				Biography:      &p.Edges.Author.Biography,
				CreatedAt:      p.Edges.Author.CreatedAt,
				Id:             p.Edges.Author.ID,
				Name:           p.Edges.Author.Name,
				Nickname:       p.Edges.Author.Nickname,
				UpdatedAt:      p.Edges.Author.UpdatedAt,
			},
			CreatedAt:      p.CreatedAt,
			Favorited:      favorited,
			FavoritesCount: favoritesCount,
			Id:             p.ID,
			Text:           &p.Text,
			UpdatedAt:      p.UpdatedAt,
		})
	}

	return posts, nextCursor, nil
}
