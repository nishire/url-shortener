package utility

import (
	"log"
	"url-shortener/internal/app/app_constants"
	"url-shortener/pkg/constants"

	"github.com/spf13/viper"
)

type IRedisKeyGenerator interface {
	GenerateRedisKeyForUrl(string) string
}

type RedisKeyGenerator struct{}

func NewRedisKeyGenerator() *RedisKeyGenerator {
	return &RedisKeyGenerator{}
}

func (u RedisKeyGenerator) GenerateRedisKeyForUrl(url string) string {
	log.Println("Creating redis key for url")
	env := viper.GetString(constants.Environment)
	var key = env + app_constants.UNDERSCORE_SEPARATOR + app_constants.SERVICE_NAME_KEY + app_constants.COLON_SEPERATOR + constants.Inhouse + app_constants.UNDERSCORE_SEPARATOR + url
	return key
}
