package container

import (
	"sync"
	"url-shortener/internal/app/controller"
	"url-shortener/internal/app/factories"
	"url-shortener/internal/app/utility"
	"url-shortener/internal/app/utility/cache"
)

var (
	dependencyInjectorInstance *dependencyInjector
	containerOnce              sync.Once
)

type IServiceContainer interface {
	InjectDependencies() *controller.UrlController
}

type dependencyInjector struct{}

func (c *dependencyInjector) InjectDependencies() *controller.UrlController {
	redisKeyGen := utility.NewRedisKeyGenerator()
	redisClient := cache.GetRedisClientImp()
	base62Encoder := utility.NewBase62Encoder()

	urlShortenerFactory := factories.NewUrlShortenerFactory(redisClient, redisKeyGen, base62Encoder)

	urlController := controller.NewUrlController(urlShortenerFactory)
	return urlController
}

func ServiceContainer() IServiceContainer {
	if dependencyInjectorInstance == nil {
		containerOnce.Do(func() {
			dependencyInjectorInstance = &dependencyInjector{}
		})
	}
	return dependencyInjectorInstance
}
