package cachefx

import (
	"context"
	"time"

	ttlcache "github.com/jellydator/ttlcache/v3"
)

// MemoryCache memory cache
type MemoryCache struct {
	cache *ttlcache.Cache[string, []byte]
}

// NewMemoryCache new memory cache
func NewMemoryCache() *MemoryCache {
	cache := ttlcache.New[string, []byte]()
	return &MemoryCache{
		cache: cache,
	}
}

func (c *MemoryCache) Get(ctx context.Context, key string) ([]byte, error) {
	item := c.cache.Get(key)
	if item != nil {
		return item.Value(), nil
	}
	return nil, ErrNotFound
}

func (c *MemoryCache) Set(ctx context.Context, key string, value []byte,
	expiration time.Duration) error {
	c.cache.Set(key, value, expiration)

	return nil
}

func (c *MemoryCache) Delete(ctx context.Context, key string) error {
	c.cache.Delete(key)
	return nil
}
