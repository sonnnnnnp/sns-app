package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/internal/pkg/errors"
	"github.com/sonnnnnnp/reverie/pkg/ws"
)

func (c Controller) Stream(ctx echo.Context) error {
	hub := ctxhelper.GetWebSocketHub(ctx.Request().Context())

	conn, err := hub.Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return errors.ErrWebsocketProtocolRequired
	}

	client := ws.NewClient(ctxhelper.GetUserID(ctx.Request().Context()), conn)
	hub.RegisterChan <- client

	defer func() {
		hub.UnregisterChan <- client
	}()

	go client.Ping()

	for {
		var msg ws.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		if err := func(msg ws.Message) error {
			switch msg.Type {
			case "subscribe":
				return c.streamUsecase.OnSubscribe(ctx.Request().Context(), client, msg)
			case "unsubscribe":
				return c.streamUsecase.OnUnsubscribe(ctx.Request().Context(), client, msg)
			default:
				return fmt.Errorf("unspported message type: %s", msg.Type)
			}
		}(msg); err != nil {
			client.Send(ws.Message{
				Type: "error",
				Body: fmt.Sprintf("%v", err),
			})
		}
	}

	return nil
}