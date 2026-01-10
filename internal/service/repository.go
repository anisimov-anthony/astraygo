package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO: think about precision
// TODO: think about types in pg
// TODO: use DTO
type ObjectInfo struct {
	ID        int64     `json:"id" redis:"id"`
	Latitude  float64   `json:"latitude" redis:"latitude"`
	Longitude float64   `json:"longitude" redis:"longitude"`
	Time      time.Time `json:"time" redis:"time"`
}

type Repository interface {
	GetAllIDs() (*[]int64, error)
	GetObjectByID(ID int64) (*ObjectInfo, error)
	UpdateObjectLocation(object *ObjectInfo) (*ObjectInfo, error)
}

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{pool: pool}
}

func (pgrepo *PostgresRepo) GetObjectByID(ID int64) (*ObjectInfo, error) {

	sql := `
			SELECT object_id, latitude, longitude, time 
			FROM objects
			WHERE object_id = $1
	`

	var object = ObjectInfo{}

	// TODO: think about context
	row := pgrepo.pool.QueryRow(context.Background(), sql, ID)
	err := row.Scan(
		&object.ID,
		&object.Latitude,
		&object.Longitude,
		&object.Time,
	)
	if err != nil {
		return nil, err
	}

	return &object, nil
}

// TODO: use PG index
func (pgrepo *PostgresRepo) GetAllIDs() (*[]int64, error) {

	sql := `
		SELECT DISTINCT object_id
		FROM objects
	`

	var IDs []int64

	// TODO: think about context
	rows, err := pgrepo.pool.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ID int64

		err := rows.Scan(&ID)
		if err != nil {
			return nil, err
		}

		IDs = append(IDs, ID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &IDs, nil
}

// TODO: use PG index
func (pgrepo *PostgresRepo) UpdateObjectLocation(object *ObjectInfo) (*ObjectInfo, error) {

	sql := `
		INSERT INTO objects (object_id, latitude, longitude, time)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (object_id) 
		DO UPDATE SET 
			latitude = EXCLUDED.latitude,
			longitude = EXCLUDED.longitude,
			time = EXCLUDED.time;
	`

	// TODO: think about context
	_, err := pgrepo.pool.Exec(context.Background(), sql, object.ID, object.Latitude, object.Longitude, object.Time)
	if err != nil {
		return nil, err
	}

	return object, nil
}
