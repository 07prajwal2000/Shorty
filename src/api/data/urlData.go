package data

import (
	data "shorty/data/postgresData"
	"shorty/models"

	"github.com/jackc/pgx/v5/pgtype"
)

type IUrlRepository interface {
	AddUrl(url *models.Url) error
	GetUrlById(id string) (string, error)
}

type UrlPostgresRepository struct{}

func (u UrlPostgresRepository) AddUrl(url *models.Url) error {
	_, err := postgresDb.CreateUrl(dbContext, data.CreateUrlParams{
		Originalurl: url.Original,
		Hash:        url.Hash,
		CreatedAt:   pgtype.Int8{Int64: url.CreatedAt},
		Owner:       pgtype.Int8{Int64: int64(url.UserID)},
	})
	return err
}

func (u UrlPostgresRepository) GetUrlById(id string) (string, error) {
	response, err := postgresDb.GetUrlByHash(dbContext, id)
	return response.Originalurl, err
}
