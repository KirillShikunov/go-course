package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	get := r.client.Get(ctx, key)
	return get.Val(), get.Err()
}

func (r *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client}
}
