package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionOptions struct {
	ConnString      string
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
	MaxConns        int32
	MinConns        int32
}

func Open(opt *ConnectionOptions) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"%s?pool_max_conn_lifetime=%s&pool_max_conn_idle_time=%s&pool_max_conns=%d&pool_min_conns=%d",
			opt.ConnString,
			opt.MaxConnLifetime,
			opt.MaxConnIdleTime,
			opt.MaxConns,
			opt.MinConns,
		),
	)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
