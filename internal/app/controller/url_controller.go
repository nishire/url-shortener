package controller

import (
	"fmt"
	"net/http"
	"url-shortener/internal/app/factories"
	"url-shortener/pkg/constants"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	urlShortenerFactory factories.IUrlShortenerFactoryInterface
}

func NewUrlController(UrlShortenerFactory factories.IUrlShortenerFactoryInterface) *UrlController {
	return &UrlController{
		urlShortenerFactory: UrlShortenerFactory,
	}
}

func (u UrlController) FetchShortUrl() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		shortenerName := c.Param(constants.UrlShortenerNameParamKey)
		urlShortService := u.urlShortenerFactory.GetUrlShortenerByName(shortenerName)
		if urlShortService == nil {
			fmt.Printf("Invalid Url Shortener name", shortenerName)
			c.JSON(http.StatusBadRequest, constants.InvalidUrlShortenerNameError)
			return
		}
		url := c.Param("url")
		respData, err := urlShortService.FetchShortUrl(url)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		resp, ok := respData.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, respData)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
	return fn
}
