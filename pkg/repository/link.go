package repository

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

type LinkRepository interface {
	Get(ctx context.Context, key string) string
	Set(ctx context.Context, key string, value string, exp int) string
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

func (r *LinkRepo) Get(ctx context.Context, key string) string {}

func (r *LinkRepo) Set(ctx context.Context, key string, value string, exp int) string {}
