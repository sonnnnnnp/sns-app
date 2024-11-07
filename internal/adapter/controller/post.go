package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) GetPostByID(ctx echo.Context, id uuid.UUID) error {
	p, err := c.postUsecase.GetPostByID(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, p)
}

func (c Controller) CreatePost(ctx echo.Context) error {
	var body oapi.CreatePostJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	p, err := c.postUsecase.CreatePost(ctx.Request().Context(), &body)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, p)
}

func (c Controller) FavoritePost(ctx echo.Context) error {
	var body oapi.FavoritePostJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	err := c.postUsecase.CreatePostFavorite(ctx.Request().Context(), body.PostId)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, nil)
}

func (c Controller) UnfavoritePost(ctx echo.Context) error {
	var body oapi.UnfavoritePostJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	err := c.postUsecase.DeletePostFavorite(ctx.Request().Context(), body.PostId)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, nil)
}

func (c Controller) GetPostFavorites(ctx echo.Context, params oapi.GetPostFavoritesParams) error {
	favorites, err := c.postUsecase.GetPostFavorites(ctx.Request().Context(), params.PostId)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, favorites)
}

func (c Controller) GetTimeline(ctx echo.Context, params oapi.GetTimelineParams) error {
	posts, nextCursor, err := c.postUsecase.GetTimeline(ctx.Request().Context(), &params)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &oapi.Timeline{
		Posts:      posts,
		NextCursor: nextCursor,
	})
}
