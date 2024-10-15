package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/internal/tools/ws"
)

func (c Controller) Stream(ctx echo.Context) error {
	upgrader := ctxhelper.GetUpgrader(ctx.Request().Context())

	conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return errors.ErrWebsocketProtocolRequired
	}
	defer conn.Close()

	// TODO: add user id to client struct
	client := &ws.Client{Conn: conn}

	var room = ws.Room{}
	room.Add(client)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			ctx.Logger().Error(err)
			break
		}

		room.Publish(msg)
	}

	return nil
}
