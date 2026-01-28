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

	pool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		panic(err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		panic(err)
	}

	enablePostGIS := `CREATE EXTENSION IF NOT EXISTS postgis;`
	if _, err := pool.Exec(context.Background(), enablePostGIS); err != nil {
		panic(err)
	}

	createTableObjects := `
			CREATE TABLE IF NOT EXISTS objects (
				object_id BIGINT PRIMARY KEY,
				status BOOLEAN NOT NULL,
				location GEOMETRY(Point, 4326) NOT NULL,
				time TIMESTAMPTZ NOT NULL
			);
	`
	if _, err := pool.Exec(context.Background(), createTableObjects); err != nil {
		panic(err)
	}

	createIndexObjectsLocation := `
		CREATE INDEX IF NOT EXISTS objects_location_idx
		ON objects USING GIST (location);
	`
	if _, err := pool.Exec(context.Background(), createIndexObjectsLocation); err != nil {
		panic(err)
	}

	return pool
}
