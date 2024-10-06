package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) GetTimeline(ctx echo.Context, params oapi.GetTimelineParams) error {
	entPosts, nextCursor, err := c.timelineUsecase.GetTimeline(ctx.Request().Context(), &params)
	if err != nil {
		return err
	}

	// DTO などの変換専用のロジックが必要かもしれない
	oapiPosts := make([]oapi.Post, len(entPosts))
	for i, entPost := range entPosts {
		oapiPosts[i] = oapi.Post{
			Id: entPost.ID,
			Author: oapi.User{
				AvatarImageUrl: &entPost.Edges.Author.AvatarImageURL,
				Biography:      &entPost.Edges.Author.Biography,
				BannerImageUrl: &entPost.Edges.Author.BannerImageURL,
				CreatedAt:      entPost.Edges.Author.CreatedAt,
				Nickname:       entPost.Edges.Author.Nickname,
				Id:             entPost.Edges.Author.ID,
				UpdatedAt:      entPost.Edges.Author.UpdatedAt,
				Name:           entPost.Edges.Author.Name,
			},
			Text:      &entPost.Text,
			UpdatedAt: entPost.UpdatedAt,
			CreatedAt: entPost.CreatedAt,
		}
	}

	return c.json(ctx, http.StatusOK, &oapi.Timeline{
		Posts:      oapiPosts,
		NextCursor: nextCursor,
	})
}
