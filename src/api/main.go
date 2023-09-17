package main

import (
	"fmt"
	"os"
	"shorty/controllers"
	"shorty/initializers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	const port = 5556
	url := fmt.Sprintf("localhost:%d", port)
	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	initializers.Initialize()

	os.Setenv("DOMAIN", fmt.Sprintf("http://%s/", url))

	if len(os.Args) > 0 {
		prod := os.Args[0]
		if prod == "true" {
			fmt.Println("Running in production")
		}
	}

	controllers.MapUrlController(app)

	app.Logger.Fatal(app.Start(url))
}
