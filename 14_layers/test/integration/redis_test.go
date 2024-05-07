package integration

import (
	cache "14_layers/internal/cache"
	"14_layers/internal/config"
	_ "14_layers/internal/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestRedisCache(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Env("REDIS_HOST"), config.Env("REDIS_PORT")),
	})

	key := "testKey"
	value := "testValue"
	ctx := context.Background()
	expireAt := time.Duration(1) * time.Minute

	redisCache := cache.NewRedisCache(client)
	redisCache.Set(ctx, key, value, expireAt)

	retrievedValue := redisCache.Get(ctx, key)

	if retrievedValue != value {
		t.Errorf("Expected %s but got %s", value, retrievedValue)
	}
}
