package data

import (
	"errors"
	"shorty/models"
	"time"
)

type UrlData struct {
	urlsMap map[string]*models.Url
}

var UrlsData *UrlData

func InitializeUrlData() {
	UrlsData = new(UrlData)
	UrlsData.urlsMap = make(map[string]*models.Url)
}

func (urlData *UrlData) AddUrl(url *models.Url) error {
	_, exists := urlData.urlsMap[url.Hash]
	if exists {
		return errors.New("key already exists")
	}
	url.CreatedAt = time.Now().Unix()
	urlData.urlsMap[url.Hash] = url
	return nil
}

func (urlData *UrlData) GetUrlById(id string) (*models.Url, error) {
	value, exists := urlData.urlsMap[id]
	if !exists {
		return nil, errors.New("url not found")
	}
	return value, nil
}
