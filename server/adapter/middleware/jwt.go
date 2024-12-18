package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/server/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/server/pkg/errors"
	"github.com/sonnnnnnp/reverie/server/pkg/jwter"
	"github.com/sonnnnnnp/reverie/server/pkg/utils"
)

func JWT(excludePaths []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// 認証が不必要なパスの除外
			if utils.IsExcludedPath(ctx.Path(), excludePaths) {
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

			cfg := ctxhelper.GetConfig(ctx.Request().Context())

			// 認証トークンを検証
			if uID, err := jwter.Verify(token, "access", []byte(cfg.JWTSecret)); err == nil {
				c := ctxhelper.SetUserID(ctx.Request().Context(), uID)
				ctx.SetRequest(ctx.Request().WithContext(c))
			} else {
				return err
			}

			// 認証トークンをコンテキストに挿入
			c := ctxhelper.SetAccessToken(ctx.Request().Context(), token)
			ctx.SetRequest(ctx.Request().WithContext(c))

			return next(ctx)
		}
	}
}
