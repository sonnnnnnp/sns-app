package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	authorize_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/authorize"
	post_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/post"
	stream_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/stream"
	user_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/user"
)

type Response struct {
	OK   bool        `json:"ok"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Controller struct {
	authUsecase   *authorize_usecase.AuthorizeUsecase
	postUsecase   *post_usecase.PostUsecase
	streamUsecase *stream_usecase.StreamUsecase
	userUsecase   *user_usecase.UserUsecase
}

func New(
	authUsecase *authorize_usecase.AuthorizeUsecase,
	postUsecase *post_usecase.PostUsecase,
	streamUsecase *stream_usecase.StreamUsecase,
	userUsecase *user_usecase.UserUsecase,
) *Controller {
	return &Controller{
		authUsecase:   authUsecase,
		postUsecase:   postUsecase,
		streamUsecase: streamUsecase,
		userUsecase:   userUsecase,
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
