package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ICacheClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type CacheClient struct {
	cc *redis.Client
}

func NewCacheClient() (*CacheClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	// TODO: implement a better way to handle this, retries, etc

	pong, err := client.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}

	fmt.Println(pong)

	return &CacheClient{
		cc: client,
	}, nil
}

func (c CacheClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := c.cc.Set(ctx, key, value, expiration).Err()

	if err != nil {
		return err
	}

	return nil
}

func (c CacheClient) Get(ctx context.Context, key string) (string, error) {
	keyValue, err := c.cc.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	fmt.Println(keyValue)

	return keyValue, nil

}
