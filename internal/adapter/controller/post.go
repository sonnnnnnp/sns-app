package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) CreatePost(ctx echo.Context) error {
	var body oapi.CreatePostJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	p, err := c.postUsecase.CreatePost(ctx.Request().Context(), &body)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &oapi.Post{
		Id: p.ID,
		Author: oapi.User{
			AvatarUrl:   p.Edges.Author.AvatarURL,
			Biography:   p.Edges.Author.Biography,
			CoverUrl:    p.Edges.Author.CoverURL,
			CreatedAt:   p.Edges.Author.CreatedAt,
			DisplayName: p.Edges.Author.DisplayName,
			Id:          p.Edges.Author.ID,
			UpdatedAt:   p.Edges.Author.UpdatedAt,
			Username:    p.Edges.Author.Username,
		},
		Content:   &p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	})
}
