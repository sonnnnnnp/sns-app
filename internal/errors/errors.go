package errors

import "errors"

var (
	// auth
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrInvalidTokenScope = errors.New("invalid token scope")

	// user
	ErrUserNotFound         = errors.New("user not found")
	ErrCannotBlockSelf      = errors.New("cannot block self")
	ErrUserAlreadyBlocking  = errors.New("user already blocking")
	ErrCannotFollowSelf     = errors.New("cannot follow self")
	ErrUserAlreadyFollowing = errors.New("user already following")

	// post
	ErrPostNotFound         = errors.New("post not found")
	ErrPostAreadlyFavorited = errors.New("post already favorited")
	ErrPostFavoriteNotFound = errors.New("post favorite not found")

	// other
	ErrCursorNotFound            = errors.New("cursor not found")
	ErrWebsocketProtocolRequired = errors.New("websocket protocol required")
)
