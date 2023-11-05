package controllers

import (
	"fmt"
	"net/http"
	"shorty/models"
	"shorty/services"

	"github.com/labstack/echo/v4"
)

var urlService services.IUrlServices

func MapUrlController(app *echo.Echo) {
	urlService = services.CreateUrlServices()

	urlGroup := app.Group("/api/v1/urls/")
	urlGroup.GET("redirect/:id", getUrlByIdRoute)
	urlGroup.POST("generate", generateUrlRoute)
}

func generateUrlRoute(c echo.Context) error {
	urlDto := &models.GenerateUrlDto{}
	if err := c.Bind(urlDto); err != nil {
		return c.JSON(http.StatusBadRequest, &models.GenerateUrlResponse{
			StatusCode: 400,
			Message:    "Invalid JSON data format",
		})
	}
	response := urlService.GenerateUrl(urlDto)
	return c.JSON(response.StatusCode, response)
}

func getUrlByIdRoute(c echo.Context) error {
	id := c.Param("id")
	if len(id) != services.HASH_SIZE {
		return c.JSON(http.StatusBadRequest, &models.GetUrlByIdResponse{
			Message:    fmt.Sprintf("Invalid URL ID: %s", id),
			StatusCode: http.StatusBadRequest,
		})
	}
	urlResponse := urlService.GetUrlById(id)
	if urlResponse.StatusCode != http.StatusOK {
		return c.JSON(http.StatusNotFound, urlResponse)
	}

	return c.JSON(http.StatusOK, urlResponse)
}
