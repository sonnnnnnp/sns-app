package errors

import "errors"

var (
	// auth
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrInvalidTokenScope = errors.New("invalid token scope")

	// user
	ErrUserNotFound            = errors.New("user not found")
	ErrCannotBlockYourself     = errors.New("cannot block yourself")
	ErrUserAlreadyBlocking     = errors.New("user already blocking")
	ErrUserBlocking            = errors.New("blocking the user")
	ErrUserBlockedBy           = errors.New("blocked by the user")
	ErrUserBlockingOrBlockedBy = errors.New("blocking or blocked by the user")
	ErrCannotFollowYourself    = errors.New("cannot follow yourself")
	ErrUserAlreadyFollowing    = errors.New("user already following")

	// post
	ErrPostNotFound         = errors.New("post not found")
	ErrPostAreadlyFavorited = errors.New("post already favorited")
	ErrPostFavoriteNotFound = errors.New("post favorite not found")

	// call
	ErrCallNotFound       = errors.New("call not found")
	ErrNotCallHost        = errors.New("not call host")
	ErrAlreadyJoinedACall = errors.New("already joined a call")
	ErrCallAlreadyEnded   = errors.New("call already ended")

	// other
	ErrCursorNotFound            = errors.New("cursor not found")
	ErrWebsocketProtocolRequired = errors.New("websocket protocol required")
	ErrInsufficientParameters    = errors.New("insufficient parameters")
)
