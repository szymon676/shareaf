package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore() *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisStore{
		client: rdb,
	}
}

func (rs *RedisStore) Set() error {
	err := rs.client.Set(ctx, "wow", "wow", time.Second*30).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rs *RedisStore) Get() error {
	result, err := rs.client.Get(ctx, "wow").Result()
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
