package main

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	SavePaste(Paste) error
	RetrievePaste(name string) (any, error)
	DeletePaste(name string) error
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

func (rs *RedisStore) SavePaste(paste Paste) error {
	err := rs.client.Set(ctx, paste.Name, paste.Data, time.Second*paste.Time).Err()
	if err != nil {
		return err
	}

	log.Print("saved paste under name:", paste.Name)
	return nil
}

func (rs *RedisStore) RetrievePaste(name string) (any, error) {
	result, err := rs.client.Get(ctx, name).Result()
	if err != nil {
		return "", err
	}

	log.Print("retieved paste under name:", name)
	return result, nil
}

func (rs *RedisStore) DeletePaste(name string) error {
	err := rs.client.Del(ctx, name).Err()
	if err != nil {
		return err
	}
	return nil
}
