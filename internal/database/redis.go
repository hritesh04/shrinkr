package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type cache struct {
	client *redis.Client
}

func InitCache() *cache {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &cache{
		client: client,
	}
}

func (c *cache) Get(shorten string) (string, error) {
	return "", nil
}
func (c *cache) Set(ctx context.Context, key string, value string, time time.Duration) error {
	return c.client.Set(ctx, key, value, time).Err()
}
