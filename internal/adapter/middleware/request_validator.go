package middleware

import (
	"context"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
)

func RequestValidator(swagger *openapi3.T) echo.MiddlewareFunc {
	return echomiddleware.OapiRequestValidatorWithOptions(
		swagger,
		&echomiddleware.Options{
			Options: openapi3filter.Options{
				ExcludeRequestBody:    false,
				ExcludeResponseBody:   false,
				IncludeResponseStatus: true,
				MultiError:            true,
				AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {
					return nil
				},
			},
			MultiErrorHandler: func(muerr openapi3.MultiError) *echo.HTTPError {
				return &echo.HTTPError{
					Code:     http.StatusBadRequest,
					Message:  muerr.Error(),
					Internal: muerr,
				}
			},
		},
	)
}
