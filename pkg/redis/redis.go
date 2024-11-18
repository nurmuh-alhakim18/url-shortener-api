package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	log.Println("Connected to redis")
	return &RedisClient{Client: client}
}

func (r *RedisClient) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockRedisClient) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	args := m.Called(ctx, key, value, ttl)
	return args.Error(0)
}
