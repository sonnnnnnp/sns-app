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
	ErrCannotBlockYourself  = errors.New("cannot block yourself")
	ErrUserAlreadyBlocking  = errors.New("user already blocking")
	ErrUserBlockedByYou     = errors.New("you are blocking the user")
	ErrUserBlockingYou      = errors.New("user blocking you")
	ErrCannotFollowYourself = errors.New("cannot follow yourself")
	ErrUserAlreadyFollowing = errors.New("user already following")

	// post
	ErrPostNotFound         = errors.New("post not found")
	ErrPostAreadlyFavorited = errors.New("post already favorited")
	ErrPostFavoriteNotFound = errors.New("post favorite not found")

	// other
	ErrCursorNotFound            = errors.New("cursor not found")
	ErrWebsocketProtocolRequired = errors.New("websocket protocol required")
)
