package cache

import (
	"context"
	"time"

	"github.com/dock-tech/notes-api/internal/config/properties"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/redis/go-redis/v9"
)

type cacheClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type cache struct {
	clientSet cacheClient
	clientGet cacheClient
}

func (c cache) Get(ctx context.Context, key string) ([]byte, error) {
	return c.clientGet.Get(ctx, properties.GetCachePrefix()+key).Bytes()
}
func (c cache) Set(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	return c.clientSet.Set(ctx, properties.GetCachePrefix()+key, value, expiration).Err()
}

func NewCache(clientSet, clientGet cacheClient) interfaces.Cache {
	return &cache{clientSet: clientSet, clientGet: clientGet}
}
