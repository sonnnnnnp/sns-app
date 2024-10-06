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
				AvatarUrl:   entPost.Edges.Author.AvatarURL,
				Biography:   entPost.Edges.Author.Biography,
				CoverUrl:    entPost.Edges.Author.CoverURL,
				CreatedAt:   entPost.Edges.Author.CreatedAt,
				DisplayName: entPost.Edges.Author.DisplayName,
				Id:          entPost.Edges.Author.ID,
				UpdatedAt:   entPost.Edges.Author.UpdatedAt,
				Username:    entPost.Edges.Author.Username,
			},
			Content:   &entPost.Content,
			UpdatedAt: entPost.UpdatedAt,
			CreatedAt: entPost.CreatedAt,
		}
	}

	return c.json(ctx, http.StatusOK, &oapi.Timeline{
		Posts:      oapiPosts,
		NextCursor: nextCursor,
	})
}
