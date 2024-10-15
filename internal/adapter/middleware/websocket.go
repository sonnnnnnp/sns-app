package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/internal/tools/ws"
)

func WebSocketMiddleware(h *ws.Hub) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c := ctxhelper.SetWebSocketHub(ctx.Request().Context(), h)
			ctx.SetRequest(ctx.Request().WithContext(c))
			return next(ctx)
		}
	}
}
