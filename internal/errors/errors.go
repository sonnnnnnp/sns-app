package errors

import "errors"

var (
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrInvalidTokenScope = errors.New("invalid token scope")

	ErrUserNotFound     = errors.New("user not found")
	ErrCannotFollowSelf = errors.New("cannot follow self")

	ErrPostAreadlyFavorited = errors.New("post already favorited")
	ErrPostFavoriteNotFound = errors.New("post favorite not found")

	ErrWebsocketProtocolRequired = errors.New("websocket protocol required")
)
