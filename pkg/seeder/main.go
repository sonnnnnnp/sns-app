package seeder

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

type Seeder struct {
	pool    *pgxpool.Pool
	queries *db.Queries
}

func New(pool *pgxpool.Pool) *Seeder {
	return &Seeder{
		pool:    pool,
		queries: db.New(pool),
	}
}

func (s *Seeder) Run() error {
	if err := s.seedUsers(); err != nil {
		return err
	}

	return nil
}

func ptr(s string) *string { return &s }
