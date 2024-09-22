package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) AuthorizeWithLine(ctx echo.Context, params oapi.AuthorizeWithLineParams) error {
	auth, err := c.authUsecase.AuthorizeWithLine(ctx.Request().Context(), params.Code)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &auth)
}

func (c Controller) RefreshAuthorization(ctx echo.Context) error {
	var body *oapi.RefreshAuthorizationJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	auth, err := c.authUsecase.RefreshAuthorization(ctx.Request().Context(), body)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &auth)
}
