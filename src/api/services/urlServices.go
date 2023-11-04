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

type IUrlServices interface {
	GenerateUrl(dto *models.GenerateUrlDto) *models.GenerateUrlResponse
	GetUrlById(id string) *models.GetUrlByIdResponse
}

type UrlServices struct {
	UrlRepository data.IUrlRepository
}

func CreateUrlServices() IUrlServices {
	return UrlServices{
		UrlRepository: data.CreateUrlRepository(),
	}
}

func (u UrlServices) GenerateUrl(dto *models.GenerateUrlDto) *models.GenerateUrlResponse {
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
	err := u.UrlRepository.AddUrl(url)
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

func (u UrlServices) GetUrlById(id string) *models.GetUrlByIdResponse {
	url, err := u.UrlRepository.GetUrlById(id)

	if err != nil || url == "" {
		return &models.GetUrlByIdResponse{
			Message:    err.Error(),
			StatusCode: http.StatusNotFound,
		}
	}
	return &models.GetUrlByIdResponse{
		OriginalUrl: url,
		StatusCode:  http.StatusOK,
		Message:     "",
	}
}

func createShortUrl(id string) string {
	domain := os.Getenv("DOMAIN")
	return fmt.Sprintf("%s%s", domain, id)
}
