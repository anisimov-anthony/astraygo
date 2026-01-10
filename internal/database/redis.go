package database

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {

	connString, ok := os.LookupEnv("REDIS_CONN_STR")
	if !ok {
		panic(fmt.Errorf("env REDIS_CONN_STR not set"))
	}

	opt, err := redis.ParseURL(connString)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	// TODO: Think about context
	ctx := context.Background()
	err = client.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	return client
}
