package container

import (
	"go-helpers/cache"
	"sync"
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
	urlShortenerFactory := factories.NewUrlShortenerFactory(redisKeyGen, redisClient)

	urlController := controller.NewUrlController(urlShortenerFactory)
	return urlController
}

// ServiceContainer : ServiceContainer
func ServiceContainer() IServiceContainer {
	if dependencyInjectorInstance == nil {
		containerOnce.Do(func() {
			dependencyInjectorInstance = &dependencyInjector{}
		})
	}
	return dependencyInjectorInstance
}
