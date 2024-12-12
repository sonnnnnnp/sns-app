package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/internal/server/http/adapter/api"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/authorize"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/call"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/call_timeline"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/post"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/stream"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/timeline"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/user"
)

type Response struct {
	OK   bool        `json:"ok"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Controller struct {
	authUsecase         *authorize.AuthorizeUsecase
	callUsecase         *call.CallUsecase
	callTimelineUsecase *call_timeline.CallTimelineUsecase
	postUsecase         *post.PostUsecase
	streamUsecase       *stream.StreamUsecase
	timelineUsecase     *timeline.TimelineUsecase
	userUsecase         *user.UserUsecase
}

func New(
	authUsecase *authorize.AuthorizeUsecase,
	callUsecase *call.CallUsecase,
	callTimelineUsecase *call_timeline.CallTimelineUsecase,
	postUsecase *post.PostUsecase,
	streamUsecase *stream.StreamUsecase,
	timelineUsecase *timeline.TimelineUsecase,
	userUsecase *user.UserUsecase,
) *Controller {
	return &Controller{
		authUsecase:         authUsecase,
		callUsecase:         callUsecase,
		callTimelineUsecase: callTimelineUsecase,
		postUsecase:         postUsecase,
		streamUsecase:       streamUsecase,
		timelineUsecase:     timelineUsecase,
		userUsecase:         userUsecase,
	}
}

func (c *Controller) json(ctx echo.Context, code int, data interface{}) error {
	return ctx.JSON(http.StatusOK, &Response{
		Code: code,
		OK:   code == http.StatusOK,
		Data: data,
	})
}

var _ api.ServerInterface = (*Controller)(nil)
