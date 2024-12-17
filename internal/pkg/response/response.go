package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	OK   bool        `json:"ok"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func JSON(ctx echo.Context, code int, data interface{}) error {
	return ctx.JSON(http.StatusOK, &Response{
		Code: code,
		OK:   code == http.StatusOK,
		Data: data,
	})
}
