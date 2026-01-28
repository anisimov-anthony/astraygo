package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ObjectInfo struct {
	ID        int64     `json:"id" redis:"id"`
	Status    bool      `json:"status" redis:"status"`
	Latitude  float64   `json:"latitude" redis:"latitude"`
	Longitude float64   `json:"longitude" redis:"longitude"`
	Time      time.Time `json:"time" redis:"time"`
}

type Repository interface {
	GetActiveIDs() ([]int64, error)

	GetAllObjects() ([]ObjectInfo, error)
	GetObjectsByStatus(status bool) ([]ObjectInfo, error)
	GetObjectByID(ID int64) (*ObjectInfo, error)

	UpdateObject(object *ObjectInfo) (*ObjectInfo, error)
}

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{pool: pool}
}

// TODO: think about context
func (pgrepo *PostgresRepo) GetActiveIDs() ([]int64, error) {
	sql := `
		SELECT DISTINCT object_id
		FROM objects
		WHERE status = true
	`

	rows, err := pgrepo.pool.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activeIDs []int64
	for rows.Next() {
		var activeID int64

		err := rows.Scan(
			&activeID,
		)
		if err != nil {
			return nil, err
		}

		activeIDs = append(activeIDs, activeID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return activeIDs, nil
}

// TODO: think about context
func (pgrepo *PostgresRepo) GetAllObjects() ([]ObjectInfo, error) {
	sql := `
		SELECT object_id,
			   status,
			   ST_Y(location) as latitude,
			   ST_X(location) as longitude,
			   time
		FROM objects
	`

	rows, err := pgrepo.pool.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var objects []ObjectInfo
	for rows.Next() {
		var object ObjectInfo
		err := rows.Scan(
			&object.ID,
			&object.Status,
			&object.Latitude,
			&object.Longitude,
			&object.Time,
		)

		if err != nil {
			return nil, err
		}

		objects = append(objects, object)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return objects, nil
}

// TODO: think about context
func (pgrepo *PostgresRepo) GetObjectsByStatus(status bool) ([]ObjectInfo, error) {
	sql := `
		SELECT object_id,
			   status,
			   ST_Y(location) as latitude,
			   ST_X(location) as longitude,
			   time
		FROM objects
		WHERE status = $1
	`

	rows, err := pgrepo.pool.Query(context.Background(), sql, status)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var objects []ObjectInfo
	for rows.Next() {
		var object ObjectInfo
		err := rows.Scan(
			&object.ID,
			&object.Status,
			&object.Latitude,
			&object.Longitude,
			&object.Time,
		)
		if err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return objects, nil
}

// TODO: think about context
func (pgrepo *PostgresRepo) GetObjectByID(ID int64) (*ObjectInfo, error) {

	sql := `
		SELECT object_id,
			   status,
			   ST_Y(location) as latitude,
			   ST_X(location) as longitude,
			   time
		FROM objects
		WHERE object_id = $1
	`

	var object = ObjectInfo{}

	row := pgrepo.pool.QueryRow(context.Background(), sql, ID)
	err := row.Scan(
		&object.ID,
		&object.Status,
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
// TODO: think about context
// TODO: need other method for adding new not updating old
func (pgrepo *PostgresRepo) UpdateObject(object *ObjectInfo) (*ObjectInfo, error) {

	sql := `
		INSERT INTO objects (object_id, location, time, status)
		VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3), 4326), $4, $5)
		ON CONFLICT (object_id)
		DO UPDATE SET
			location = EXCLUDED.location,
			time = EXCLUDED.time,
			status = EXCLUDED.status;
	`

	// ST_MakePoint takes (longitude, latitude) order
	_, err := pgrepo.pool.Exec(context.Background(), sql, object.ID, object.Longitude, object.Latitude, object.Time, true)
	if err != nil {
		return nil, err
	}

	return object, nil
}
