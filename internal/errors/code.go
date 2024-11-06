package errors

import (
	"errors"
	"net/http"
)

var errorCodeMap = map[error]int{
	// API errors
	ErrUnauthorized:      -1000,
	ErrInvalidToken:      -1001,
	ErrTokenExpired:      -1002,
	ErrInvalidTokenScope: -1003,

	ErrUserNotFound:     -1100,
	ErrCannotFollowSelf: -1101,

	ErrPostAreadlyFavorited: -1201,
	ErrPostFavoriteNotFound: -1202,

	// Network errors
	ErrWebsocketProtocolRequired: -10000,
}

func getErrorCode(err error) int {
	for e, code := range errorCodeMap {
		if errors.Is(err, e) {
			return code
		}
	}
	return http.StatusInternalServerError
}
