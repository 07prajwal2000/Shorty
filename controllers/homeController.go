package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: Work in progress
func MapHomeController(app *echo.Echo) {
	group := app.Group("/")
	group.GET("", homeRoute)
	group.GET("*", notFoundRoute)
}

func homeRoute(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")
}

func notFoundRoute(c echo.Context) error {
	return c.Render(http.StatusNotFound, "notfound", "")
}