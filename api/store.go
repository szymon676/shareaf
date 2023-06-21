package main

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	SavePaste(name string, data string) error
	RetrievePaste(name string) (string, error)
}

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(options RediStoreOptions) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
	})

	return &RedisStore{
		client: client,
	}
}

var ctx = context.Background()

func (rs *RedisStore) SavePaste(name string, data string) error {
	err := rs.client.Set(ctx, name, data, time.Second*86400).Err()
	if err != nil {
		return err
	}

	log.Print("saved paste under name:", name)
	return nil
}

func (rs *RedisStore) RetrievePaste(name string) (string, error) {
	result, err := rs.client.Get(ctx, name).Result()
	if err != nil {
		return "", err
	}

	log.Print("retieved paste under name:", name)
	return result, nil
}
