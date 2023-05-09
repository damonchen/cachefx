package cachefx

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// RedisCache redis cache
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache new redis cache
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	cmd := r.client.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		return nil, fmt.Errorf("get cache from redis %s", err)
	}
	return cmd.Bytes()

}

func (r *RedisCache) Set(ctx context.Context, key string, value []byte,
	expiration time.Duration) error {
	cmd := r.client.Set(ctx, key, value, expiration)
	if err := cmd.Err(); err != nil {
		return fmt.Errorf("set cache from redis %s", err)
	}
	return nil
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	cmd := r.client.Del(ctx, key)
	if err := cmd.Err(); err != nil {
		return fmt.Errorf("delete cache from redis %s", err)
	}
	return nil
}
