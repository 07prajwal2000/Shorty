package models

type Url struct {
	Original  string `json:"originalUrl"`
	ShortUrl  string `json:"shortUrl"`
	CreatedAt int64  `json:"createdAt"`
	Hash      string `json:"hash"`
	Id        int    `json:"id"`
	UserID    int    `json:"userId"`
}

type GenerateUrlResponse struct {
	Original   string `json:"originalUrl"`
	Modified   string `json:"modifiedUrl"`
	ShortUrl   string `json:"shortUrl"`
	Hash       string `json:"hash"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type GetUrlByIdResponse struct {
	OriginalUrl string `json:"originalUrl"`
	StatusCode  int    `json:"statusCode"`
	Message     string `json:"message"`
}

type GenerateUrlDto struct {
	OriginalUrl string `json:"originalUrl"`
}
