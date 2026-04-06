package postgre

import "github.com/jackc/pgx/v5/pgxpool"

type Postgre struct{
	pool *pgxpool.Pool
}

