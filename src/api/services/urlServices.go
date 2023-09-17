package services

import (
	"fmt"
	"net/http"
	"os"
	"shorty/data"
	"shorty/models"
	"shorty/utils"
	"strings"
)

const HASH_SIZE = 8

func GenerateUrl(dto *models.GenerateUrlDto) *models.GenerateUrlResponse {
	randomId := utils.GenerateRandomHash(HASH_SIZE)
	shortUrl := createShortUrl(randomId)
	url := &models.Url {
		Original: dto.OriginalUrl,
		Hash: randomId,
		ShortUrl: shortUrl,
	}
	err := data.UrlsData.AddUrl(url)

	if err != nil {
		return &models.GenerateUrlResponse{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return &models.GenerateUrlResponse{
		Original: dto.OriginalUrl,
		ShortUrl: shortUrl,
		Hash: randomId,
		StatusCode: http.StatusOK,
		Message: "Successfully created",
	}
}

func GetUrlById(id string) *models.GetUrlByIdResponse {
	url, err := data.UrlsData.GetUrlById(id)
	
	if err != nil || url == nil {
		return &models.GetUrlByIdResponse{
			Message: err.Error(),
			StatusCode: http.StatusNotFound,
		}
	}
	
	if !strings.HasPrefix(url.Original, "http://") {
		url.Original = fmt.Sprintf("http://%s", url.Original)
	}
	
	return &models.GetUrlByIdResponse{
		OriginalUrl: url.Original,
		StatusCode: http.StatusOK,
		Message: "",
	}
}

func createShortUrl(id string) string {
	domain := os.Getenv("DOMAIN")
	return fmt.Sprintf("%s%s", domain, id)
}