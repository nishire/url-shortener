package utility

import (
	"context"
	"log"
	"url-shortener/internal/app/app_constants"
	"url-shortener/pkg/constants"

	"github.com/spf13/viper"
)

type IRedisKeyGenerator interface {
	GenerateRedisKeyForUrl(ctx context.Context) string
}

type RedisKeyGenerator struct{}

func NewRedisKeyGenerator() *RedisKeyGenerator {
	return &RedisKeyGenerator{}
}

func (u RedisKeyGenerator) GenerateRedisKeyForUrl(ctx context.Context) string {
	log.Println("Creating redis key for url")
	env := viper.GetString(constants.Environment)
	var key = env + app_constants.UNDERSCORE_SEPARATOR + app_constants.SERVICE_NAME_KEY + app_constants.COLON_SEPERATOR + constants.InHouse + app_constants.UNDERSCORE_SEPARATOR // need to add unique id in the end
	return key
}
