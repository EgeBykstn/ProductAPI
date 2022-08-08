package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var RC *redis.Client

func ConnectRedis() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost: 6379",
		Password: "",
		DB:       0,
	})
	_, err := redis.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal(err)
	}

	return redis
}
