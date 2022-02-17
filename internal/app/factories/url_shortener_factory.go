package factories

import (
	"url-shortener/go-helpers/cache"
	"url-shortener/internal/app/interfaces"
	"url-shortener/internal/app/services"
	"url-shortener/internal/app/utility"
	"url-shortener/pkg/constants"
)

type IUrlShortenerFactoryInterface interface {
	GetUrlShortenerByName(string) interfaces.IUrlShortenerInterface
}
type UrlShortenerFactory struct {
	redisCache  cache.IRedisClient
	redisKeyGen utility.IRedisKeyGenerator
}

func NewUrlShortenerFactory(RedisCache cache.IRedisClient, RedisKeyGen utility.IRedisKeyGenerator) IUrlShortenerFactoryInterface {
	return &UrlShortenerFactory{
		redisCache:  RedisCache,
		redisKeyGen: RedisKeyGen,
	}
}

func (u UrlShortenerFactory) GetUrlShortenerByName(name string) interfaces.IUrlShortenerInterface {
	switch name {
	case constants.Inhouse:
		return services.NewInhouseService(u.redisCache, u.redisKeyGen)
	default:
		return nil
	}
}
