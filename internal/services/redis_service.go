package services

import (
	"BudgetApp/internal/configs"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisService() *RedisService {
	client, _ := configs.ConnectionToRedis()

	return &RedisService{
		client: client,
		ctx:    context.Background(),
	}
}

// PingRedis checks the Redis connection
func (s *RedisService) PingRedis() error {
	pong, err := s.client.Ping(s.ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println("Redis connected:", pong)
	return nil
}

// SetKey sets a key-value pair in Redis
func (s *RedisService) SetKey(key string, value string) error {
	return s.client.Set(s.ctx, key, value, 0).Err()
}

// GetKey retrieves a value by key from Redis
func (s *RedisService) GetKey(key string) (string, error) {
	return s.client.Get(s.ctx, key).Result()
}
