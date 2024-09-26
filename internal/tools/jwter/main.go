package jwter

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
)

func Verify(tokenString string, scope string, secret []byte) (uID uuid.UUID, err error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return uuid.Nil, internal_errors.ErrInvalidToken
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			return uuid.Nil, internal_errors.ErrTokenExpired
		default:
			return uuid.Nil, internal_errors.ErrInvalidToken
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if scopeClaim, ok := claims["scope"].(string); ok {
			if scopeClaim != scope {
				return uuid.Nil, internal_errors.ErrInvalidTokenScope
			}
		}

		if sub, ok := claims["sub"].(string); ok {
			uID, err := uuid.Parse(sub)
			if err != nil {
				return uuid.Nil, internal_errors.ErrInvalidToken
			}
			return uID, nil
		}
		return uuid.Nil, internal_errors.ErrInvalidToken
	}

	return uuid.Nil, internal_errors.ErrInvalidToken
}
