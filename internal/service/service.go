package service

import (
	"errors"
	"sync"
)

type AstrayService struct {
	mu          sync.RWMutex
	existantIDs *map[int64]bool
	pgRepo      *PostgresRepo
	redisCache  *RedisCache
}

// TODO: improve func signature to general
func AstrayServiceInit(postgres *PostgresRepo, redis *RedisCache) *AstrayService {
	// Fetch existant IDs into memory before start
	listIDs, err := postgres.GetAllIDs()
	if err != nil {
		panic(err)
	}

	existantIDs := make(map[int64]bool)
	for _, ID := range *listIDs {
		existantIDs[ID] = true
	}

	return &AstrayService{
		existantIDs: &existantIDs,
		pgRepo:      postgres,
		redisCache:  redis,
	}
}

func (s *AstrayService) GetAllIDs() *map[int64]bool {
	return s.existantIDs
}

// FIXME: add mutex for existantIDs map
func (s *AstrayService) GetObjectByID(ID int64) (*ObjectInfo, error) {
	if _, ok := (*s.existantIDs)[ID]; !ok {
		return nil, errors.New("non existant ID\n")
	}

	var object *ObjectInfo

	object, err := s.redisCache.GetObjectByID(ID)
	if err != nil {
		object, err = s.pgRepo.GetObjectByID(ID)
		if err != nil {
			return nil, errors.New("no object matches ID\n")

		}
	}

	return object, nil
}

// TODO: do concurent updates
func (s *AstrayService) UpdateObjectLocation(object *ObjectInfo) (*ObjectInfo, error) {
	var newObject *ObjectInfo

	newObject, err := s.pgRepo.UpdateObjectLocation(object)
	if err != nil {
		return nil, errors.New("Failed to update object in postgres\n")
	}

	s.mu.Lock()
	(*s.existantIDs)[object.ID] = true
	s.mu.Unlock()

	_, err = s.redisCache.UpdateObjectLocation(object)
	if err != nil {
		return nil, errors.New("Failed to update object in redis\n")
	}

	return newObject, nil
}
