package cache

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var once sync.Once
var redisClient *RedisClientImp

type IRedisClient interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, ttl time.Duration) (string, error)
	Del(keys ...string) (int64, error)
	ObtainLock(key string) error
}

type RedisClientImp struct {
	RedisClient *redis.Client
}

// GetRedisClientImp : Returns new redis client after initializing and validating the connection to the redis distributed cache
func GetRedisClientImp() *RedisClientImp {
	return redisClient
}

func (u RedisClientImp) Get(key string) (string, error) {
	return u.RedisClient.Get(context.Background(), key).Result()
}

func (u RedisClientImp) Set(key string, value interface{}, ttl time.Duration) (string, error) {
	return u.RedisClient.Set(context.Background(), key, value, ttl).Result()
}

func (u RedisClientImp) GetRedisClient() *RedisClientImp {
	return redisClient
}
