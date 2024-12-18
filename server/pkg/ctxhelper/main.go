package ctxhelper

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/reverie/server/config"
	"github.com/sonnnnnnp/reverie/server/pkg/ws"
)

// context キーの衝突を防止するため非公開独自キーを定義
type configKey struct{}
type websocketHubKey struct{}
type accessTokenKey struct{}
type userIDkey struct{}

func SetConfig(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, configKey{}, cfg)
}

func GetConfig(ctx context.Context) *config.Config {
	cfg, _ := ctx.Value(configKey{}).(*config.Config)
	return cfg
}

func SetWebSocketHub(ctx context.Context, h *ws.Hub) context.Context {
	return context.WithValue(ctx, websocketHubKey{}, h)
}

func GetWebSocketHub(ctx context.Context) *ws.Hub {
	h, _ := ctx.Value(websocketHubKey{}).(*ws.Hub)
	return h
}

func SetAccessToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, accessTokenKey{}, token)
}

func GetAccessToken(ctx context.Context) string {
	token, _ := ctx.Value(accessTokenKey{}).(string)
	return token
}

func SetUserID(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDkey{}, id)
}

func GetUserID(ctx context.Context) uuid.UUID {
	token, _ := ctx.Value(userIDkey{}).(uuid.UUID)
	return token
}
