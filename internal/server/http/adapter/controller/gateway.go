package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/internal/pkg/response"
)

func (c Controller) GetGatewayURI(ctx echo.Context) error {
	uri, err := c.gatewayUsecase.GetGatewayURI(ctx.Request().Context())
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, uri)
}
