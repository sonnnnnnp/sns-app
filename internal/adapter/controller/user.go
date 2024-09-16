package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) GetUser(ctx echo.Context, userId string) error {
	u, err := c.db.User.Create().SetUsername(userId).SetDisplayName(userId).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to create a user")
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
