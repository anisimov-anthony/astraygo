package service

import (
	"errors"
)

type AstrayService struct {
	pgRepo     *PostgresRepo
	redisCache *RedisCache
}

// TODO: improve func signature to general
func AstrayServiceInit(postgres *PostgresRepo, redis *RedisCache) *AstrayService {
	return &AstrayService{
		pgRepo:     postgres,
		redisCache: redis,
	}
}

func (s *AstrayService) GetObjects(status *bool) ([]ObjectInfo, error) {
	var objects []ObjectInfo
	var err error

	if status == nil {
		objects, err = s.pgRepo.GetAllObjects()
	} else {
		objects, err = s.pgRepo.GetObjectsByStatus(*status)
	}

	if err != nil {
		return nil, errors.New("failed to get objects from database")
	}
	return objects, nil
}

func (s *AstrayService) GetObjectByID(ID int64) (*ObjectInfo, error) {
	var object *ObjectInfo

	object, err := s.redisCache.GetObjectByID(ID)
	if err != nil {
		object, err = s.pgRepo.GetObjectByID(ID)
		if err != nil {
			return nil, errors.New("no object matches ID")

		}
	}

	return object, nil
}

// TODO: do concurent updates
func (s *AstrayService) UpdateObjectLocation(object *ObjectInfo) (*ObjectInfo, error) {
	var newObject *ObjectInfo

	newObject, err := s.pgRepo.UpdateObject(object)
	if err != nil {
		return nil, errors.New("Failed to update object in postgres")
	}

	_, err = s.redisCache.UpdateObject(object)
	if err != nil {
		return nil, errors.New("Failed to update object in redis")
	}

	return newObject, nil
}
