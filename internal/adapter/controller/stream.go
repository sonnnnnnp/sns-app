package controller

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/internal/tools/ws"
)

func (c Controller) Stream(ctx echo.Context) error {
	h := ctxhelper.GetWebSocketHub(ctx.Request().Context())

	conn, err := h.Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return errors.ErrWebsocketProtocolRequired
	}

	client := ws.NewClient(ctxhelper.GetUserID(ctx.Request().Context()), conn)
	h.Register <- client

	defer func() {
		h.Unregister <- client
	}()

	for {
		var msg ws.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		log.Printf("%v", msg)

		switch msg.Type {
		case "subscribe":
			c.streamUsecase.OnSubscribe(ctx.Request().Context(), h, client, msg)
		case "unsubscribe":
			c.streamUsecase.OnUnsubscribe(ctx.Request().Context(), h, client, msg)
		default:
			h.Broadcast <- ws.Message{
				Type: "error",
				Body: "unknown operation type",
			}
		}
	}

	return nil
}
