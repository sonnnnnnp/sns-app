package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/internal/adapter/api"
)

// user

func (c Controller) GetSelf(ctx echo.Context) error {
	u, err := c.userUsecase.GetSelf(ctx.Request().Context())
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, u)
}

func (c Controller) GetUser(ctx echo.Context, params api.GetUserParams) error {
	u, err := c.userUsecase.GetUser(ctx.Request().Context(), params)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, u)
}

func (c Controller) UpdateUser(ctx echo.Context) error {
	var body api.UpdateUserJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	u, err := c.userUsecase.UpdateUser(ctx.Request().Context(), &body)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, u)
}

// block

func (c Controller) BlockUser(ctx echo.Context, uID uuid.UUID) error {
	err := c.userUsecase.BlockUser(ctx.Request().Context(), uID)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, make(map[string]interface{}))
}

func (c Controller) GetUserBlocking(ctx echo.Context) error {
	users, err := c.userUsecase.GetUserBlocking(ctx.Request().Context())
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &api.Users{
		Users: users,
	})
}

func (c Controller) UnblockUser(ctx echo.Context, uID uuid.UUID) error {
	err := c.userUsecase.UnblockUser(ctx.Request().Context(), uID)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, make(map[string]interface{}))
}

// follow

func (c Controller) GetUserFollowing(ctx echo.Context, uID uuid.UUID) error {
	users, err := c.userUsecase.GetUserFollowing(ctx.Request().Context(), uID)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &api.UserFollowers{
		Users: users,
	})
}

func (c Controller) GetUserFollowers(ctx echo.Context, uID uuid.UUID) error {
	users, err := c.userUsecase.GetUserFollowers(ctx.Request().Context(), uID)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &api.UserFollowers{
		Users: users,
	})
}

func (c Controller) FollowUser(ctx echo.Context, uID uuid.UUID) error {
	sc, err := c.userUsecase.FollowUser(ctx.Request().Context(), uID)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, sc)
}

func (c Controller) UnfollowUser(ctx echo.Context, uID uuid.UUID) error {
	sc, err := c.userUsecase.UnfollowUser(ctx.Request().Context(), uID)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, sc)
}

func (c Controller) RemoveUserFromFollowers(ctx echo.Context, uID uuid.UUID) error {
	sc, err := c.userUsecase.RemoveUserFromFollowers(ctx.Request().Context(), uID)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, sc)
}
