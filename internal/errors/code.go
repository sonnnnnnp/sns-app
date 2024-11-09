package errors

import (
	"errors"
	"net/http"
)

var errorCodeMap = map[error]int{
	// auth
	ErrUnauthorized:      -1000,
	ErrInvalidToken:      -1001,
	ErrTokenExpired:      -1002,
	ErrInvalidTokenScope: -1003,

	// user
	ErrUserNotFound:         -1100,
	ErrCannotBlockSelf:      -1101,
	ErrUserAlreadyBlocking:  -1102,
	ErrCannotFollowSelf:     -1103,
	ErrUserAlreadyFollowing: -1104,

	// post
	ErrPostNotFound:         -1200,
	ErrPostAreadlyFavorited: -1201,
	ErrPostFavoriteNotFound: -1202,

	// other
	ErrCursorNotFound:            -10000,
	ErrWebsocketProtocolRequired: -10001,
}

func getErrorCode(err error) int {
	for e, code := range errorCodeMap {
		if errors.Is(err, e) {
			return code
		}
	}
	return http.StatusInternalServerError
}
