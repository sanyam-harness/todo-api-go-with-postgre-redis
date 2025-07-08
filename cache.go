package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Default Redis port
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	return &RedisClient{client: rdb}
}

func (r *RedisClient) SetTodosCache(todos []*Todo) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, "todos_cache", data, 10*time.Second).Err()
}

func (r *RedisClient) GetTodosCache() ([]*Todo, error) {
	val, err := r.client.Get(ctx, "todos_cache").Result()
	if err == redis.Nil {
		return nil, nil // Cache miss
	} else if err != nil {
		return nil, err
	}

	var todos []*Todo
	if err := json.Unmarshal([]byte(val), &todos); err != nil {
		return nil, err
	}
	return todos, nil
}
