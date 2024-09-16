package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) AuthorizeWithLine(ctx echo.Context, params oapi.AuthorizeWithLineParams) error {
	return c.json(ctx, http.StatusOK, &oapi.Authorization{
		AccessToken:  "1234567890",
		RefreshToken: "abcdefghijklmnopqrstuvwxyz",
		UserId:       params.Code,
	})
}

func (c Controller) RefreshAuthorization(ctx echo.Context) error {
	return c.json(ctx, http.StatusOK, &oapi.Authorization{
		AccessToken:  "1234567890",
		RefreshToken: "abcdefghijklmnopqrstuvwxyz",
		UserId:       "or4p90.fo0qg4",
	})
}
