# url-shortener

API Info :
1. Here, you need to use the api /url-shortener/api/v1/fetchShortUrl to fetch a short url for the provided  long url.
2. The request body looks like :
    {
    "original_url": "anywebsitename.com"
    }
3. The response is a shortUrl in string format in case of status code 200 i.e successful response.

Architecture :
1. Here we have implemented a clean architecture and have used factory pattern.
2. Considering we can have multiple providers that provide shortening of url, factory instance for each can be fetched with function GetUrlShortenerByName. Currently, we only have inhouse url shortener.
3. Inside the controller function we first fetch provider name i.e the provider who does url-shortening for us as I mentioned we can have multiple providers.
   a. Fetch provider name i.e shortenerName (here we have "inhouse")
   b. Fetch interface with all method signatures for that particular provider.
   c. Use these methods further as required in the controller.