package postgre

import "github.com/jackc/pgx/v5/pgxpool"

type Postgre struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Postgre {
	return &Postgre{pool: pool}
}
