package initializers

import (
	"shorty/controllers"
	"shorty/data"

	"github.com/labstack/echo/v4"
)

func Initialize() {
	data.InitializeUrlData()
}

func InitializeControllers(app *echo.Echo) {
	controllers.MapUrlController(app)
	controllers.MapTokensController(app)
}
