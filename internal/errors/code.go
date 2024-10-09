package errors

import (
	"errors"
	"net/http"
)

var errorCodeMap = map[error]int{
	ErrUnauthorized:      -1000,
	ErrInvalidToken:      -1001,
	ErrTokenExpired:      -1002,
	ErrInvalidTokenScope: -1003,

	ErrUserNotFound:     -2000,
	ErrCannotFollowSelf: -2010,

	ErrPostAreadlyFavorited: -3001,
}

func getErrorCode(err error) int {
	for e, code := range errorCodeMap {
		if errors.Is(err, e) {
			return code
		}
	}
	return http.StatusInternalServerError
}
