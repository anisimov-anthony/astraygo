package service

import (
	"context"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	GetObjectByID(ID int64) (*ObjectInfo, error)
	UpdateObjectLocation(object *ObjectInfo) (*ObjectInfo, error)
	WarmUp(repo Repository)
}

type RedisCache struct {
	redisClient *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{redisClient: client}
}

func (r *RedisCache) WarmUp(repo Repository) {
	listIDs, err := repo.GetAllIDs()
	if err != nil {
		panic(err)
	}

	for _, ID := range *listIDs {
		object, err := repo.GetObjectByID(ID)
		if err != nil {
			panic(err)
		}

		result, err := r.UpdateObjectLocation(object)
		if err != nil {
			panic(err)
		}

		// TODO: use logger (slog, etc.)
		log.Printf("Warmed Up object: %v\n", result)
	}
}

func (r *RedisCache) GetObjectByID(ID int64) (*ObjectInfo, error) {
	var object ObjectInfo

	// TODO: think about context
	err := r.redisClient.HGetAll(context.Background(), strconv.FormatInt(ID, 10)).Scan(&object)

	if err != nil {
		return nil, err
	}

	return &object, nil
}

func (r *RedisCache) UpdateObjectLocation(object *ObjectInfo) (*ObjectInfo, error) {

	// TODO: think about context
	err := r.redisClient.HSet(context.Background(), strconv.FormatInt(object.ID, 10), object).Err()

	if err != nil {
		return nil, err
	}

	return object, nil
}
