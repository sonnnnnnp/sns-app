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
			AvatarImageUrl: &p.Edges.Author.AvatarImageURL,
			Biography:      &p.Edges.Author.Biography,
			BannerImageUrl: &p.Edges.Author.BannerImageURL,
			CreatedAt:      p.Edges.Author.CreatedAt,
			Nickname:       p.Edges.Author.Nickname,
			Id:             p.Edges.Author.ID,
			UpdatedAt:      p.Edges.Author.UpdatedAt,
			Name:           p.Edges.Author.Name,
		},
		Text:      &p.Text,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	})
}
