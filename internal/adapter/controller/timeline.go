package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
)

func (c Controller) GetTimeline(ctx echo.Context, params api.GetTimelineParams) error {
	posts, nextCursor, err := c.timelineUsecase.GetTimeline(ctx.Request().Context(), &params)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &api.Timeline{
		Posts:      posts,
		NextCursor: nextCursor,
	})
}

func (c Controller) GetFollowingTimeline(ctx echo.Context, params api.GetFollowingTimelineParams) error {
	posts, nextCursor, err := c.timelineUsecase.GetFollowingTimeline(ctx.Request().Context(), &params)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &api.Timeline{
		Posts:      posts,
		NextCursor: nextCursor,
	})
}

func (c Controller) GetUserTimeline(ctx echo.Context, uID uuid.UUID, params api.GetUserTimelineParams) error {
	posts, nextCursor, err := c.timelineUsecase.GetUserTimeline(ctx.Request().Context(), uID, &params)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, &api.Timeline{
		Posts:      posts,
		NextCursor: nextCursor,
	})
}
