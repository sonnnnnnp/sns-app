package config

import "github.com/caarlos0/env/v11"

type Config struct {
	DBHost     string `env:"DB_HOST" envDefault:"db"`
	DBUser     string `env:"DB_USER" envDefault:"user"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"password"`
	DBName     string `env:"DB_NAME" envDefault:"db"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
