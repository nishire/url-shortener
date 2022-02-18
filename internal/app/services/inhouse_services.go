package services

import (
	"fmt"
	"hash/fnv"
	"net/url"
	"strings"
	"url-shortener/go-helpers/cache"
	"url-shortener/internal/app/interfaces"
	"url-shortener/internal/app/utility"
	"url-shortener/pkg/constants"

	"github.com/golang/glog"
	"github.com/spf13/viper"
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

func (i inhouse) FetchShortUrl(inputUrl interface{}) (interface{}, error) {
	var shortUrl string
	var err error
	originalUrl := inputUrl.(string)

	// formatting url
	originalUrl = " " + originalUrl
	originalUrl = strings.Trim(originalUrl, "https://")
	originalUrl = strings.Trim(originalUrl, "www.")
	originalUrl = strings.TrimSpace(originalUrl)

	originalUrl = fmt.Sprintf("https://www.%s", originalUrl)

	// validating url
	_, err = url.ParseRequestURI(originalUrl)
	if err != nil {
		glog.Error("Url Parse Failed", err)
		return nil, err
	}

	// get us a redis key that can be associated to a short url
	key := i.redisKeyGenerator.GenerateRedisKeyForUrl(originalUrl)

	// check if key already exists
	isExists := i.redisCache.Exists(key)
	if isExists {
		shortUrl, err = i.redisCache.Get(key)
		if err != nil {
			glog.Error("Fetch data from redis/cache failed", err)
			return shortUrl, err
		}
	}

	// create unique hash
	var shortUrlUniqueId uint32
	algorithm := fnv.New32a()
	algorithm.Write([]byte(originalUrl))
	shortUrlUniqueId = algorithm.Sum32()

	// convert that unique key to url path
	shortUrl = fmt.Sprintf("%s/%d", viper.GetString(constants.UrlShortenerBaseUrl), shortUrlUniqueId)

	// save to cache memory
	_, err = i.redisCache.Set(key, shortUrl, viper.GetDuration("redis.timeout"))
	if err != nil {
		glog.Error("Save shortUrl in memory failed")
		return shortUrl, err
	}
	return shortUrl, nil
}
