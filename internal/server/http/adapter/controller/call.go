package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/internal/server/http/adapter/api"
)

func (c Controller) CreateCall(ctx echo.Context) error {
	var body api.CreateCallJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	room, err := c.callUsecase.CreateCall(ctx.Request().Context(), body)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &room)
}

func (c Controller) EndCall(ctx echo.Context, cID uuid.UUID) error {
	if err := c.callUsecase.EndCall(ctx.Request().Context(), cID); err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, make(map[string]interface{}))
}
