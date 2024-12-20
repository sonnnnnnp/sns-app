package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/pkg/response"
)

func (c Controller) GetTimeline(ctx echo.Context, params api.GetTimelineParams) error {
	posts, nextCursor, err := c.timelineUsecase.GetTimeline(ctx.Request().Context(), &params)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, &api.PostTimeline{
		Posts:      posts,
		NextCursor: nextCursor,
	})
}

func (c Controller) GetFollowingTimeline(ctx echo.Context, params api.GetFollowingTimelineParams) error {
	posts, nextCursor, err := c.timelineUsecase.GetFollowingTimeline(ctx.Request().Context(), &params)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, &api.PostTimeline{
		Posts:      posts,
		NextCursor: nextCursor,
	})
}

func (c Controller) GetUserTimeline(ctx echo.Context, uID uuid.UUID, params api.GetUserTimelineParams) error {
	posts, nextCursor, err := c.timelineUsecase.GetUserTimeline(ctx.Request().Context(), uID, &params)
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, &api.PostTimeline{
		Posts:      posts,
		NextCursor: nextCursor,
	})
}
