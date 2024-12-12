package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/config"
)

func Config(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c := ctxhelper.SetConfig(ctx.Request().Context(), cfg)
			ctx.SetRequest(ctx.Request().WithContext(c))
			return next(ctx)
		}
	}
}