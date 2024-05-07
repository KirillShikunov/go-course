package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func (r *RedisCache) Get(ctx context.Context, key string) string {
	return r.client.Get(ctx, key).Val()
}

func (r *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) {
	r.client.Set(ctx, key, value, expiration)
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client}
}
