package container

import (
	"sync"
	"url-shortener/go-helpers/cache"
	"url-shortener/internal/app/controller"
	"url-shortener/internal/app/factories"
	"url-shortener/internal/app/utility"
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
	urlShortenerFactory := factories.NewUrlShortenerFactory(redisClient, redisKeyGen)

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
