package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPostgres() *pgxpool.Pool {

	connString, ok := os.LookupEnv("PG_CONN_STR")
	if !ok {
		panic(fmt.Errorf("env PG_CONN_STR not set"))
	}

	pgConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		panic(err)
	}

	// TODO: think about context
	pool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		panic(err)
	}

	// TODO: think about context
	if err := pool.Ping(context.Background()); err != nil {
		panic(err)
	}

	// TODO: need index
	sql := `
			CREATE TABLE IF NOT EXISTS objects (
				object_id BIGINT PRIMARY KEY,
				latitude DOUBLE PRECISION NOT NULL,
				longitude DOUBLE PRECISION NOT NULL,
				time TIMESTAMPTZ NOT NULL
			);
	`

	// TODO: think about context
	if _, err := pool.Exec(context.Background(), sql); err != nil {
		panic(err)
	}

	return pool
}
