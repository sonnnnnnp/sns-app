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

	if u == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.json(ctx, http.StatusOK, &oapi.User{
		Id:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		AvatarUrl:   u.AvatarURL,
		CoverUrl:    u.CoverURL,
		Biography:   u.Biography,
		Birthdate:   u.Birthdate,
		UpdatedAt:   u.UpdatedAt,
		CreatedAt:   u.CreatedAt,
	})
}

func (c Controller) GetUserById(ctx echo.Context, id uuid.UUID) error {
	u, err := c.userUsecase.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	if u == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.json(ctx, http.StatusOK, &oapi.User{
		Id:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		AvatarUrl:   u.AvatarURL,
		CoverUrl:    u.CoverURL,
		Biography:   u.Biography,
		Birthdate:   u.Birthdate,
		UpdatedAt:   u.UpdatedAt,
		CreatedAt:   u.CreatedAt,
	})
}

func (c Controller) UpdateUser(ctx echo.Context, id uuid.UUID) error {
	var body *oapi.UpdateUserJSONBody
	if err := ctx.Bind(&body); err != nil {
		// return errors.ValidationFailed.Wrap(errors.WithStack(err), err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u, err := c.userUsecase.UpdateUser(ctx.Request().Context(), id, body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to update a user")
	}

	return c.json(ctx, http.StatusOK, &oapi.User{
		Id:          u.Id,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		AvatarUrl:   u.AvatarUrl,
		CoverUrl:    u.CoverUrl,
		Biography:   u.Biography,
		Birthdate:   u.Birthdate,
		UpdatedAt:   u.UpdatedAt,
		CreatedAt:   u.CreatedAt,
	})
}
