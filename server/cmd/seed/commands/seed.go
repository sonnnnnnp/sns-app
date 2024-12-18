package commands

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/reverie/server/infra/db"
)

type Seed struct {
	pool    *pgxpool.Pool
	queries *db.Queries
}

func New(pool *pgxpool.Pool) *Seed {
	return &Seed{
		pool:    pool,
		queries: db.New(pool),
	}
}

func (s *Seed) Run() error {
	if err := s.users(); err != nil {
		return err
	}

	return nil
}
