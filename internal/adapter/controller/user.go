package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) GetSelf(ctx echo.Context) error {
	uID := ctxhelper.GetUserID(ctx.Request().Context())

	u, err := c.userUsecase.GetUserByID(ctx.Request().Context(), uID)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, u)
}

func (c Controller) GetUserByName(ctx echo.Context, name string) error {
	u, err := c.userUsecase.GetUserByName(ctx.Request().Context(), name)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, u)
}

func (c Controller) UpdateUser(ctx echo.Context) error {
	uID := ctxhelper.GetUserID(ctx.Request().Context())

	var body oapi.UpdateUserJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	u, err := c.userUsecase.UpdateUser(ctx.Request().Context(), uID, &body)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, u)
}

func (c Controller) GetUserFollowing(ctx echo.Context, id uuid.UUID) error {
	users, err := c.userUsecase.GetUserFollowing(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &oapi.Users{
		Users: users,
	})
}

func (c Controller) GetUserFollowers(ctx echo.Context, id uuid.UUID) error {
	users, err := c.userUsecase.GetUserFollowers(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &oapi.Users{
		Users: users,
	})
}

func (c Controller) FollowUser(ctx echo.Context, id uuid.UUID) error {
	sc, err := c.userUsecase.FollowUser(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, sc)
}

func (c Controller) UnfollowUser(ctx echo.Context, id uuid.UUID) error {
	sc, err := c.userUsecase.UnfollowUser(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, sc)
}

func (c Controller) RemoveUserFromFollowers(ctx echo.Context, id uuid.UUID) error {
	sc, err := c.userUsecase.RemoveUserFromFollowers(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, sc)
}
