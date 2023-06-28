package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	SavePaste(name any, data any) error
	RetrievePaste(name any) (any, error)
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

func (rs *RedisStore) SavePaste(name any, data any) error {
	convname, ok := name.(string)
	if !ok {
		return errors.New("wrong paste name, please insert again")
	}

	err := rs.client.Set(ctx, convname, data, time.Second*86400).Err()
	if err != nil {
		return err
	}

	log.Print("saved paste under name:", name)
	return nil
}

func (rs *RedisStore) RetrievePaste(name any) (any, error) {
	convname, ok := name.(string)
	if !ok {
		return nil, errors.New("wrong paste name, please insert again")
	}

	result, err := rs.client.Get(ctx, convname).Result()
	if err != nil {
		return "", err
	}

	log.Print("retieved paste under name:", name)
	return result, nil
}
