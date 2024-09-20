package middleware

import (
	"errors"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
)

var (
	ErrInvalidJWT          = errors.New("invalid token")
	ErrInvalidJWTSignature = errors.New("signature required")
	ErrInvalidJWTID        = errors.New("id required")
	ErrExpiredJWT          = errors.New("token expired")
)

func verifyJWT(tokenString string, secret []byte) (uID uuid.UUID, err error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return uuid.Nil, ErrInvalidJWTSignature
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			return uuid.Nil, ErrExpiredJWT
		default:
			return uuid.Nil, ErrInvalidJWT
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(string); ok {
			uID, err := uuid.Parse(id)
			if err != nil {
				return uuid.Nil, ErrInvalidJWTID
			}
			return uID, nil
		}
		return uuid.Nil, ErrInvalidJWTID
	}

	return uuid.Nil, ErrInvalidJWT
}

func JWTMiddleware(excludePaths []string) echo.MiddlewareFunc {
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
				return echo.NewHTTPError(401, "unauthorized")
			}

			authParts := strings.Split(auth, " ")
			if len(authParts) != 2 || authParts[0] != "Bearer" {
				return echo.NewHTTPError(401, "invalid authorization")
			}

			token := authParts[1]
			if token == "" {
				return echo.NewHTTPError(401, "invalid token format")
			}

			// 認証トークンを検証
			cfg := ctxhelper.GetConfig(ctx.Request().Context())

			if uID, err := verifyJWT(token, []byte(cfg.JWTSecret)); err == nil {
				c := ctxhelper.SetUserID(ctx.Request().Context(), uID)
				ctx.SetRequest(ctx.Request().WithContext(c))
			} else {
				return echo.NewHTTPError(401, err.Error())
			}

			// 認証トークンを挿入
			c := ctxhelper.SetAccessToken(ctx.Request().Context(), token)
			ctx.SetRequest(ctx.Request().WithContext(c))

			return next(ctx)
		}
	}
}
