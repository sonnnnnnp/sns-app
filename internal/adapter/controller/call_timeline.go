package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
)

func (c Controller) GetCallTimeline(ctx echo.Context, params api.GetCallTimelineParams) error {
	rooms, nextCursor, err := c.callTimelineUsecase.GetCallTimeline(ctx.Request().Context(), params)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &api.CallTimeline{
		Rooms:      rooms,
		NextCursor: nextCursor,
	})
}

func (c Controller) GetFollowingCallTimeline(ctx echo.Context, params api.GetFollowingCallTimelineParams) error {
	rooms, nextCursor, err := c.callTimelineUsecase.GetFollowingCallTimeline(ctx.Request().Context(), params)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &api.CallTimeline{
		Rooms:      rooms,
		NextCursor: nextCursor,
	})
}

func (c Controller) GetUserCallTimeline(ctx echo.Context, uID uuid.UUID, params api.GetUserCallTimelineParams) error {
	rooms, nextCursor, err := c.callTimelineUsecase.GetUserCallTimeline(ctx.Request().Context(), uID, params)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &api.CallTimeline{
		Rooms:      rooms,
		NextCursor: nextCursor,
	})
}
