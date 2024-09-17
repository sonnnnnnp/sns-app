package middleware

import (
	"log"
	"slices"

	"github.com/labstack/echo/v4"
)

// TODO: implement jwt authentication
func JWTMiddleware(excludePaths []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			path := ctx.Path()

			if slices.Contains(excludePaths, path) {
				log.Printf("exclude path: %s", path)
				return next(ctx)
			}

			log.Printf("path: %s", path)

			// get jwt (access token) from request headers
			// check jwt validity
			// check jwt expiration
			// insert jwt to context

			return next(ctx)
		}
	}
}
