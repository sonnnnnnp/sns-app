package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DBDSN             string        `env:"DB_DSN" envDefault:"postgres://user:password@db:5432/db"`
	DBMaxConnLifetime time.Duration `env:"DB_MAX_CONN_LIFETIME" envDefault:"30m"`
	DBMaxConnIdleTime time.Duration `env:"DB_MAX_CONN_IDLE_TIME" envDefault:"5m"`
	DBMaxConns        int32         `env:"DB_MAX_CONNS" envDefault:"10"`
	DBMinConns        int32         `env:"DB_MIN_CONNS" envDefault:"1"`

	JWTSecret string `env:"JWT_SECRET" envDefault:"jwt_secret"`

	LineAuthRedirectURL  string `env:"LINE_AUTH_REDIRECT_URL"`
	LineAuthClientID     string `env:"LINE_AUTH_CLIENT_ID"`
	LineAuthClientSecret string `env:"LINE_AUTH_CLIENT_SECRET"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
