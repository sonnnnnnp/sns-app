package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/pkg/response"
)

func (c Controller) GetPostByID(ctx echo.Context, pID uuid.UUID) error {
	p, err := c.postUsecase.GetPostByID(ctx.Request().Context(), pID)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, p)
}

func (c Controller) CreatePost(ctx echo.Context) error {
	var body api.CreatePostJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	p, err := c.postUsecase.CreatePost(ctx.Request().Context(), &body)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, p)
}

func (c Controller) DeletePost(ctx echo.Context, pID uuid.UUID) error {
	err := c.postUsecase.DeletePost(ctx.Request().Context(), pID)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, make(map[string]interface{}))
}

func (c Controller) FavoritePost(ctx echo.Context, pID uuid.UUID) error {
	err := c.postUsecase.CreatePostFavorite(ctx.Request().Context(), pID)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, make(map[string]interface{}))
}

func (c Controller) UnfavoritePost(ctx echo.Context, pID uuid.UUID) error {
	err := c.postUsecase.DeletePostFavorite(ctx.Request().Context(), pID)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, make(map[string]interface{}))
}

func (c Controller) GetPostFavorites(ctx echo.Context, pID uuid.UUID) error {
	favorites, err := c.postUsecase.GetPostFavorites(ctx.Request().Context(), pID)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, favorites)
}
