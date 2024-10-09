package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) GetTimeline(ctx echo.Context, params oapi.GetTimelineParams) error {
	posts, nextCursor, err := c.timelineUsecase.GetTimeline(ctx.Request().Context(), &params)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &oapi.Timeline{
		Posts:      posts,
		NextCursor: nextCursor,
	})
}
