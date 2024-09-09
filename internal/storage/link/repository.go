package link

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const expiration = 0

func NewRepo(db *redis.Client) *LinkRepo {
	return &LinkRepo{
		db: db,
	}
}

type LinkRepo struct {
	db *redis.Client
}

func (r *LinkRepo) Get(ctx context.Context, key string) (string, error) {
	return r.db.Get(ctx, key).Result()
}

func (r *LinkRepo) Set(ctx context.Context, key string, value string) error {
	return r.db.Set(ctx, key, value, expiration).Err()
}
