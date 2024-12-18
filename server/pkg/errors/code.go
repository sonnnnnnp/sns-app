package errors

import (
	"errors"
	"net/http"
)

// エラーコードは6桁の整数で表す。形式は以下の通り。
// - [メインカテゴリ]
// - [サブカテゴリ]
// - [エラー番号]
//
// 例えば、120104は以下のように解釈される。
// - 12 (ユーザー)
// - 01 (ブロック)
// - 04 (特定のエラー)

var errorCodeMap = map[error]int{
	// -- common --
	ErrWebsocketProtocolRequired: -100000,
	ErrInsufficientParameters:    -100001,
	ErrCursorNotFound:            -100002,

	// -- auth --
	ErrUnauthorized:      -110000,
	ErrInvalidToken:      -110001,
	ErrTokenExpired:      -110002,
	ErrInvalidTokenScope: -110003,

	// -- users --
	ErrUserNotFound: -120000,
	// blocks
	ErrCannotBlockYourself:     -120100,
	ErrUserAlreadyBlocking:     -120101,
	ErrUserBlocking:            -120102,
	ErrUserBlockedBy:           -120103,
	ErrUserBlockingOrBlockedBy: -120104,
	// follows
	ErrCannotFollowYourself: -120200,
	ErrUserAlreadyFollowing: -120201,

	// -- posts --
	ErrPostNotFound:         -130000,
	ErrPostAreadlyFavorited: -130001,
	ErrPostFavoriteNotFound: -130002,

	// -- calls --

	ErrCallNotFound:       -140000,
	ErrNotCallHost:        -140001,
	ErrAlreadyJoinedACall: -140002,
	ErrCallAlreadyEnded:   -140003,
}

func getErrorCode(err error) int {
	for e, code := range errorCodeMap {
		if errors.Is(err, e) {
			return code
		}
	}
	return http.StatusInternalServerError
}
