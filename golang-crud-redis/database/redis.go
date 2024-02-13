package database

import (
	"fmt"
	"go-crud-redis-example/config"

	"github.com/go-redis/redis/v8"
)

func ConnectionRedisDb(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})

	fmt.Println("Connected Successfuly to the database (Redis)")

	return rdb
}
