package app

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

type Redis interface {
	RedisSet(ctx context.Context, key string, value any) error
	RedisGet(ctx context.Context, key string) string
}

func NewRedisClient(host, port, password string) Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0, // use default DB
	})

	res := client.Ping(context.Background())
	log.Println("Redis: ", res)
	return &redisClient{client}
}

func (r *redisClient) RedisSet(ctx context.Context, key string, value any) error {
	v := r.client.Set(ctx, key, value, 0)

	return v.Err()
}

func (r *redisClient) RedisGet(ctx context.Context, key string) string {
	v := r.client.Get(ctx, key)
	return v.Val()
}
