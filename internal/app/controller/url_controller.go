package controller

import (
	"net/http"
	"url-shortener/internal/app/factories"
	"url-shortener/pkg/constants"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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
		// refer README to understand what is shortenerName --> 3.a.
		shortenerName := "inhouse"
		urlShortService := u.urlShortenerFactory.GetUrlShortenerByName(shortenerName)
		if urlShortService == nil {
			glog.Error("Invalid Url Shortener name", shortenerName)
			c.JSON(http.StatusBadRequest, constants.InvalidUrlShortenerNameError)
			return
		}
		var requestData struct {
			OriginalUrl string `json:"original_url"`
		}
		err := c.BindJSON(&requestData)
		if err != nil {
			glog.Error("Bind Request Data Failed...")
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		respData, err := urlShortService.FetchShortUrl(requestData.OriginalUrl)
		if err != nil {
			glog.Error("FetchShortUrl Failed...")
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		resp, ok := respData.(string)
		if !ok {
			glog.Error("Something went wrong while type casting short url")
			c.JSON(http.StatusBadRequest, respData)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
	return fn
}
