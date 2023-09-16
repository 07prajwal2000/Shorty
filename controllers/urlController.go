package controllers

import (
	"fmt"
	"net/http"
	"os"
	"shorty/models"
	"shorty/services"

	"github.com/labstack/echo/v4"
)

var urlsMap map[string]string = make(map[string]string)
var domain = os.Getenv("DOMAIN")
const HASH_SIZE = 8

func MapUrlController(app *echo.Echo) {
	urlGroup := app.Group("/api/v1/urls/");
	urlGroup.POST("generate", generateUrlRoute)
	urlGroup.GET(":id", getUrlByIdRoute)
}

func generateUrlRoute(c echo.Context) error {
	urlDto := &models.GenerateUrlDto {}
	if err := c.Bind(urlDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	randomId := services.GenerateRandom(HASH_SIZE)
	urlsMap[randomId] = urlDto.OriginalUrl
	response := &models.UrlResponse{
		Original: urlDto.OriginalUrl,
		ShortUrl: fmt.Sprintf("%s%s", domain, randomId),
		Id: randomId,
	}
	return c.JSON(http.StatusOK, response)
}

func getUrlByIdRoute(c echo.Context) error {
	id := &struct{
		Id string `param:"id"`
	} { }
	err := c.Bind(id)

	if err != nil || len(id.Id) != HASH_SIZE {
		return c.String(http.StatusBadRequest, "Invalid URL ID")
	}
	url, contains := urlsMap[id.Id]
	if !contains {
		return c.String(http.StatusBadRequest, fmt.Sprintf("No URL found with %s this ID", id))
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}