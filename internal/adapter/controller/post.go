package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) CreatePost(ctx echo.Context) error {
	var body *oapi.CreatePostJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	p, u, err := c.postUsecase.CreatePost(ctx.Request().Context(), body)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &oapi.Post{
		Id: p.ID,
		Author: oapi.User{
			AvatarUrl:   u.AvatarURL,
			Biography:   u.Biography,
			Birthdate:   u.Birthdate,
			CoverUrl:    u.CoverURL,
			CreatedAt:   u.CreatedAt,
			DisplayName: u.DisplayName,
			Id:          u.ID,
			UpdatedAt:   u.UpdatedAt,
			Username:    u.Username,
		},
		Content:   &p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	})
}
