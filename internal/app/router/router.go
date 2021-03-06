package router

import (
	"strings"
	"url-shortener/internal/app/container"
	"url-shortener/pkg/constants"

	cors "github.com/rs/cors/wrapper/gin"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

//NewRouter :
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	trustedProxies := viper.GetString(constants.TrustedProxies)
	trustedProxiesList := strings.Split(trustedProxies, ",")
	router.SetTrustedProxies(trustedProxiesList)

	allowedOrigins := viper.GetString(constants.AllowedOrigins)
	allowedOriginsList := strings.Split(allowedOrigins, ",")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: allowedOriginsList,
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"*"},
	})

	router.Use(corsHandler)

	urlController := container.ServiceContainer().InjectDependencies()

	url_shortener := router.Group(constants.Url_Shortener)
	{
		api := url_shortener.Group(constants.API)
		{
			v1 := api.Group(constants.Version_V1)
			{
				v1.GET(constants.FetchShortUrl, urlController.FetchShortUrl())
			}
		}
	}
	return router
}
