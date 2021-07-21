package structs

import "github.com/jackc/pgx/v4/pgxpool"

type Depend struct {
	Config
	Connection
}

type Connection struct {
	Postgres *pgxpool.Pool
}
