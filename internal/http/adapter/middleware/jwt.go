package middleware

import (
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/internal/http/errors"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/internal/pkg/jwter"
)

func JWT(excludePaths []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// 認証が不必要なパスの除外
			path := ctx.Path()
			if slices.Contains(excludePaths, path) {
				return next(ctx)
			}

			// 認証トークンの存在をチェック
			auth := ctx.Request().Header.Get("Authorization")
			if auth == "" {
				return errors.ErrUnauthorized
			}

			authParts := strings.Split(auth, " ")
			if len(authParts) != 2 || authParts[0] != "Bearer" {
				return errors.ErrInvalidToken
			}

			token := authParts[1]
			if token == "" {
				return errors.ErrInvalidToken
			}

			// 認証トークンを検証
			cfg := ctxhelper.GetConfig(ctx.Request().Context())

			if uID, err := jwter.Verify(token, "access", []byte(cfg.JWTSecret)); err == nil {
				c := ctxhelper.SetUserID(ctx.Request().Context(), uID)
				ctx.SetRequest(ctx.Request().WithContext(c))
			} else {
				return err
			}

			// 認証トークンを挿入
			c := ctxhelper.SetAccessToken(ctx.Request().Context(), token)
			ctx.SetRequest(ctx.Request().WithContext(c))

			return next(ctx)
		}
	}
}
