package constants

const (
	// Service Name
	ServiceName = "url-shortener"
	Inhouse     = "inhouse"

	//Viper Keys
	AllowedOrigins = "AllowedOrigins"
	TrustedProxies = "TrustedProxies"
	Environment    = "Environment"

	// URL contexts
	Url_Shortener = "/url-shortener"
	API           = "api"
	Version_V1    = "v1"
	FetchShortUrl = "fetchShortUrl"

	//Config File Path
	ConfigFilePath = "./pkg"

	//Param Key
	UrlShortenerNameParamKey = "url-shortener-name"

	//Error Message
	InvalidUrlShortenerNameError = "Invalid UrlShortener name provided"

	// Environment
	DevEnvironment = "dev"
)
