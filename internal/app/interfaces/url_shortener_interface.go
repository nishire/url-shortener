package interfaces

type IUrlShortenerInterface interface {
	FetchShortUrl(data interface{}) (interface{}, error)
}
