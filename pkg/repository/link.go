package repository

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type LinkRepository interface {
	Get(key string) (string, error)
	Set(key string, value string, exp time.Duration) error
}

func NewLinkRepo() *LinkRepo {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	return &LinkRepo{
		rdb: rdb,
	}
}

type LinkRepo struct {
	rdb *redis.Client
}

func (r *LinkRepo) Get(key string) (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	return r.rdb.Get(ctx, key).Result()
}

func (r *LinkRepo) Set(key string, value string, exp time.Duration) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	return r.rdb.Set(ctx, key, value, exp).Err()
}
