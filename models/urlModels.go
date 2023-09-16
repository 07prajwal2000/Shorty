package models

type UrlResponse struct {
	Original string `json:"originalUrl"`
	ShortUrl string `json:"shortUrl"`
	Id       string `json:"id"`
}

type GenerateUrlDto struct {
	OriginalUrl string `json:"originalUrl"`
}