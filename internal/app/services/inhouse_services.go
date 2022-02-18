package services

import (
	"fmt"
	"net/url"
	"strings"
	"url-shortener/internal/app/interfaces"
	"url-shortener/internal/app/utility"
	"url-shortener/internal/app/utility/cache"
	"url-shortener/pkg/constants"

	"github.com/golang/glog"
	"github.com/spf13/viper"
)

var uniqueCount uint64

type inhouse struct {
	redisCache        cache.IRedisClient
	redisKeyGenerator utility.IRedisKeyGenerator
	base62Encoder     utility.IBase62Encoder
}

func NewInhouseService(RedisCache cache.IRedisClient, RedisKeyGen utility.IRedisKeyGenerator, Base62Encoder utility.IBase62Encoder) interfaces.IUrlShortenerInterface {
	return &inhouse{
		redisCache:        RedisCache,
		redisKeyGenerator: RedisKeyGen,
		base62Encoder:     Base62Encoder,
	}
}

func (i inhouse) FetchShortUrl(inputUrl interface{}) (interface{}, error) {
	var shortUrl string
	var err error
	originalUrl := inputUrl.(string)

	// formatting url
	if !(strings.Contains(originalUrl, "https://") || strings.Contains(originalUrl, "http://")) {
		if strings.Contains(originalUrl, "www.") {
			originalUrl = fmt.Sprintf("https://%s", originalUrl)
		} else {
			originalUrl = fmt.Sprintf("https://www.%s", originalUrl)
		}
	} else {
		if !strings.Contains(originalUrl, "www.") {
			originalUrl = originalUrl + " "
			originalUrl = strings.Trim(originalUrl, "https://")
			originalUrl = strings.TrimSpace(originalUrl)
			originalUrl = fmt.Sprintf("https://www.%s", originalUrl)
		}
	}

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
		return shortUrl, nil
	}

	// maintain a unique counter id each time for generating base62
	uniqueCount = uniqueCount + 1

	// convert that unique key to encoded string
	encodedString := i.base62Encoder.ConvertUniqueKeyToUrlPath(uniqueCount)

	// convert that unique key to url path
	shortUrl = fmt.Sprintf("%s/%s", viper.GetString(constants.UrlShortenerBaseUrl), encodedString)

	// save to cache memory
	_, err = i.redisCache.Set(key, shortUrl, viper.GetDuration("RedisTimeout"))
	if err != nil {
		glog.Error("Save shortUrl in memory failed")
		return shortUrl, err
	}
	return shortUrl, nil
}
