package main

import (
	redis "github.com/redis/go-redis/v9"
)

// InitRedis initializes and returns a Redis client instance.
// It connects to Redis running on localhost at the default port 6379.
func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Default Redis port
		Password: "",       // No password by default
		DB:       0,                // Default DB
	})

	return rdb
}
