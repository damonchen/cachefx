package cachefx

import (
	"context"
	"errors"
	"time"
)

var (
	// ErrNotFound not found
	ErrNotFound = errors.New("key not exists")
)

type Cache interface {
	// Get get value from cache with value
	Get(ctx context.Context, key string) ([]byte, error)
	// Set set value to the cache with key of ttl
	Set(ctx context.Context, key string, value []byte, expiration time.Duration) error
	// Delete clear item in the cache
	Delete(ctx context.Context, key string) error
}
