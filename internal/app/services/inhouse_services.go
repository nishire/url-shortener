package services

import (
	"fmt"
	"url-shortener/go-helpers/cache"
	"url-shortener/internal/app/interfaces"
	"url-shortener/internal/app/utility"
)

type inhouse struct {
	redisCache        cache.IRedisClient
	redisKeyGenerator utility.IRedisKeyGenerator
}

func NewInhouseService(RedisCache cache.IRedisClient, RedisKeyGen utility.IRedisKeyGenerator) interfaces.IUrlShortenerInterface {
	return &inhouse{
		redisCache:        RedisCache,
		redisKeyGenerator: RedisKeyGen,
	}
}

func (i inhouse) FetchShortUrl(url interface{}) (interface{}, error) {
	originalUrl := url.(string)
	fmt.Println(originalUrl)
	return nil, nil
}
