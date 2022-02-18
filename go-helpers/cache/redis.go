package cache

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var once sync.Once
var redisClient *RedisClientImp

type IRedisClient interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, ttl time.Duration) (string, error)
	Exists(key string) bool
}

type RedisClientImp struct {
	RedisClient *redis.Client
}

// GetRedisClientImp : Returns new redis client after initializing and validating the connection to the redis distributed cache
func GetRedisClientImp() *RedisClientImp {
	once.Do(func() {
		redisURL := viper.GetString("redis.url")
		client := redis.NewClient(&redis.Options{
			Addr: redisURL,
			DB:   0, // use default DB
		})

		pingResponse, err := client.Ping(context.Background()).Result()
		if err != nil {
			fmt.Printf("Error while pinging redis cluster. ErrorMessage: %s", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Pinged redis server. Response: %s", pingResponse)
		redisClient = &RedisClientImp{RedisClient: client}
	})
	return redisClient
}

func (u RedisClientImp) Get(key string) (string, error) {
	return u.RedisClient.Get(context.Background(), key).Result()
}

func (u RedisClientImp) Set(key string, value interface{}, ttl time.Duration) (string, error) {
	return u.RedisClient.Set(context.Background(), key, value, ttl).Result()
}

func (u RedisClientImp) Exists(key string) bool {
	result := u.RedisClient.Exists(context.Background(), key)
	var isAlreadyUsed bool = int(result.Val()) != 0
	return isAlreadyUsed
}

func (u RedisClientImp) GetRedisClient() *RedisClientImp {
	return redisClient
}
