package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MapHomeController(app *echo.Echo) {
	// serves the index.html to all the not-found routes, serves to the specific route, if any static files found
	staticMiddleware := middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper:    nil,
		Root:       "static",
		Index:      "index.html",
		HTML5:      true,
		Browse:     false,
		IgnoreBase: false,
		Filesystem: nil,
	})
	app.Use(staticMiddleware)
}
