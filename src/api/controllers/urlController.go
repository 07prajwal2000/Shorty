package controllers

import (
	"fmt"
	"net/http"
	"shorty/models"
	"shorty/services"

	"github.com/labstack/echo/v4"
)

func MapUrlController(app *echo.Echo) {
	urlGroup := app.Group("/api/v1/urls/");
	app.GET(":id", getUrlByIdRoute)
	urlGroup.POST("generate", generateUrlRoute)
}

func generateUrlRoute(c echo.Context) error {
	urlDto := &models.GenerateUrlDto {}
	if err := c.Bind(urlDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	response := services.GenerateUrl(urlDto)
	return c.JSON(response.StatusCode, response)
}

func getUrlByIdRoute(c echo.Context) error {
	id := c.Param("id")

	if len(id) != services.HASH_SIZE {
		return c.JSON(http.StatusBadRequest, &models.GetUrlByIdResponse{
			Message: fmt.Sprintf("Invalid URL ID: %s", id),
			StatusCode: http.StatusBadRequest,
		})
	}

	urlResponse := services.GetUrlById(id)
	if urlResponse.StatusCode != http.StatusOK {
		return c.JSON(http.StatusNotFound, urlResponse)
	}

	return c.Redirect(http.StatusPermanentRedirect, urlResponse.OriginalUrl)
}