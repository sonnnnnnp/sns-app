package ctxhelper

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/config"
)

// context キーの衝突を防止するため非公開独自キーを定義
type configKey struct{}

func SetConfig(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, configKey{}, cfg)
}

func GetConfig(ctx context.Context) *config.Config {
	cfg, _ := ctx.Value(configKey{}).(*config.Config)
	return cfg
}
