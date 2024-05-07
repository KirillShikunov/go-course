package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) string
	Set(ctx context.Context, key string, value string, expireAt time.Duration)
}
