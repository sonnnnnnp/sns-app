package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/internal/pkg/errors"
	"github.com/sonnnnnnp/reverie/internal/pkg/jwter"
	"github.com/sonnnnnnp/reverie/pkg/ws"
)

func (c Controller) Stream(ctx echo.Context) error {
	// 認証トークンの存在をチェック
	token := ctx.QueryParam("t")
	if token == "" {
		return errors.ErrInvalidToken
	}

	cfg := ctxhelper.GetConfig(ctx.Request().Context())

	// 認証トークンを検証
	uID, err := jwter.Verify(token, "gateway", []byte(cfg.JWTSecret))
	if err == nil {
		c := ctxhelper.SetUserID(ctx.Request().Context(), uID)
		ctx.SetRequest(ctx.Request().WithContext(c))
	} else {
		return err
	}

	hub := ctxhelper.GetWebSocketHub(ctx.Request().Context())

	// プロトコルのアップグレード
	conn, err := hub.Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return errors.ErrWebsocketProtocolRequired
	}

	client := ws.NewClient(uID, conn)
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
				// return c.streamUsecase.OnSubscribe(ctx.Request().Context(), client, msg)
				return nil
			case "unsubscribe":
				// return c.streamUsecase.OnUnsubscribe(ctx.Request().Context(), client, msg)
				return nil
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
