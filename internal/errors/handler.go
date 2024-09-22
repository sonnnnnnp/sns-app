package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	OK   bool        `json:"ok"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorMessage struct {
	Message interface{} `json:"message"`
}

func ErrorHandler(err error, ctx echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		ctx.JSON(http.StatusOK, &Response{
			Code: he.Code,
			OK:   false,
			Data: &ErrorMessage{
				Message: he.Message,
			},
		})
		return
	}

	code := getErrorCode(err)

	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		OK:   false,
		Data: &ErrorMessage{
			Message: err.Error(),
		},
	})
}
