package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	authorize_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/authorize"
	post_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/post"
	stream_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/stream"
	timeline_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/timeline"
	user_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/user"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type Response struct {
	OK   bool        `json:"ok"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Controller struct {
	authUsecase     *authorize_usecase.AuthorizeUsecase
	postUsecase     *post_usecase.PostUsecase
	streamUsecase   *stream_usecase.StreamUsecase
	timelineUsecase *timeline_usecase.TimelineUsecase
	userUsecase     *user_usecase.UserUsecase
}

func New(
	authUsecase *authorize_usecase.AuthorizeUsecase,
	postUsecase *post_usecase.PostUsecase,
	streamUsecase *stream_usecase.StreamUsecase,
	timelineUsecase *timeline_usecase.TimelineUsecase,
	userUsecase *user_usecase.UserUsecase,
) *Controller {
	return &Controller{
		authUsecase:     authUsecase,
		postUsecase:     postUsecase,
		streamUsecase:   streamUsecase,
		timelineUsecase: timelineUsecase,
		userUsecase:     userUsecase,
	}
}

func (c *Controller) json(ctx echo.Context, code int, data interface{}) error {
	return ctx.JSON(http.StatusOK, &Response{
		Code: code,
		OK:   code == http.StatusOK,
		Data: data,
	})
}

var _ oapi.ServerInterface = (*Controller)(nil)
