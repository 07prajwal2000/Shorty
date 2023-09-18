package services

import (
	"fmt"
	"net/http"
	"os"
	"shorty/data"
	"shorty/models"
	"shorty/utils"
)

const HASH_SIZE = 8

func GenerateUrl(dto *models.GenerateUrlDto) *models.GenerateUrlResponse {
	randomId := utils.GenerateRandomHash(HASH_SIZE)
	shortUrl := createShortUrl(randomId)
	recievedUrl := utils.AddHttpUrlPrefix(dto.OriginalUrl)

	if !utils.IsValidUrl(recievedUrl) {
		return &models.GenerateUrlResponse{
			Original:   dto.OriginalUrl,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid URL",
		}
	}

	url := &models.Url{
		Original: recievedUrl,
		Hash:     randomId,
		ShortUrl: shortUrl,
	}
	err := data.UrlsData.AddUrl(url)
	if err != nil {
		return &models.GenerateUrlResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}

	return &models.GenerateUrlResponse{
		Original:   dto.OriginalUrl,
		Modified:   recievedUrl,
		ShortUrl:   shortUrl,
		Hash:       randomId,
		StatusCode: http.StatusOK,
		Message:    "Successfully created",
	}
}

func GetUrlById(id string) *models.GetUrlByIdResponse {
	url, err := data.UrlsData.GetUrlById(id)

	if err != nil || url == nil {
		return &models.GetUrlByIdResponse{
			Message:    err.Error(),
			StatusCode: http.StatusNotFound,
		}
	}

	return &models.GetUrlByIdResponse{
		OriginalUrl: url.Original,
		StatusCode:  http.StatusOK,
		Message:     "",
	}
}

func createShortUrl(id string) string {
	domain := os.Getenv("DOMAIN")
	return fmt.Sprintf("%s%s", domain, id)
}
