package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/pkg/response"
)

func (c Controller) AuthorizeWithLine(ctx echo.Context, params api.AuthorizeWithLineParams) error {
	auth, err := c.authUsecase.AuthorizeWithLine(ctx.Request().Context(), params.Code)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, &auth)
}

func (c Controller) RefreshAuthorization(ctx echo.Context) error {
	var body api.RefreshAuthorizationJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	auth, err := c.authUsecase.RefreshAuthorization(ctx.Request().Context(), &body)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, &auth)
}

// for development only

func (c Controller) AuthorizeWithCustomID(ctx echo.Context, customID string) error {
	auth, err := c.authUsecase.AuthorizeWithCustomID(ctx.Request().Context(), customID)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, &auth)
}
