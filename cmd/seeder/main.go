package main

import (
	"log"

	"github.com/sonnnnnnp/reverie/internal/adapter/gateway/db"
	"github.com/sonnnnnnp/reverie/pkg/config"
	"github.com/sonnnnnnp/reverie/pkg/seeder"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed initializing config: %v", err)
	}

	pool, err := db.Open(&db.ConnectionOptions{
		ConnString:      cfg.DBDSN,
		MaxConnLifetime: cfg.DBMaxConnLifetime,
		MaxConnIdleTime: cfg.DBMaxConnIdleTime,
		MaxConns:        cfg.DBMaxConns,
		MinConns:        cfg.DBMinConns,
	})
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer pool.Close()

	s := seeder.New(pool)
	if err := s.Run(); err != nil {
		log.Fatalf("failed while seeding: %v", err)
	}
}
