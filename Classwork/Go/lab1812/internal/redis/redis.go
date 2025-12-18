package redisStorage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

var ctx context.Context

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx = context.Background()
}

func SetKey(key string, val any, expTime time.Duration) error {
	err := client.Set(ctx, key, val, expTime).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func GetKey(key string) (val any, err error) {
	val, err = client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val, err
}

func GetRedis() *redis.Client {
	return client
}
