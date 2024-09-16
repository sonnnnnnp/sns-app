package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/sonnnnnnp/sns-app/internal/usecase/user"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type Response struct {
	OK   bool        `json:"ok"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Controller struct {
	userUsecase *usecase.UserUsecase
}

func New(userUsecase *usecase.UserUsecase) *Controller {
	return &Controller{
		userUsecase: userUsecase,
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
