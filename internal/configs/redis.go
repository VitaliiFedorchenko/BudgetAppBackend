package configs

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var ctx = context.Background()

func ConnectionToRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "", // No password by default
		DB:       0,  // Default DB
	})

	// Test the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
