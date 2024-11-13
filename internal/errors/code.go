package errors

import (
	"errors"
	"net/http"
)

// エラーコードは6桁の整数で表します。形式は以下の通りです。
// - [メインカテゴリ]
// - [サブカテゴリ]
// - [エラー番号]
//
// 例えば、120104 は以下のように解釈されます。
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
	ErrCannotBlockYourself:     -120101,
	ErrUserAlreadyBlocking:     -120102,
	ErrUserBlocking:            -120103,
	ErrUserBlockedBy:           -120104,
	ErrUserBlockingOrBlockedBy: -120105,
	// follows
	ErrCannotFollowYourself: -120206,
	ErrUserAlreadyFollowing: -120207,

	// -- posts --
	ErrPostNotFound:         -130000,
	ErrPostAreadlyFavorited: -130001,
	ErrPostFavoriteNotFound: -130002,
}

func getErrorCode(err error) int {
	for e, code := range errorCodeMap {
		if errors.Is(err, e) {
			return code
		}
	}
	return http.StatusInternalServerError
}
