package repository

import (
	"context"

	"github.com/go-redis/redis/v9"
)

func NewRedis(ctx context.Context, config Config) *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       0,
	})

	return redis
}
