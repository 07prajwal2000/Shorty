package main

import (
	"fmt"
	"io"
	"os"
	"shorty/controllers"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main()  {
	const port = 5556
	url := fmt.Sprintf("localhost:%d", port)
	app := echo.New()
	app.Renderer = &Template {
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	os.Setenv("DOMAIN", fmt.Sprintf("http://%s/", url))

	if len(os.Args) > 0 {
		prod := os.Args[0]
		if prod == "true" {
			fmt.Println("Running in production")
		}
	}
	
	controllers.MapHomeController(app)
	controllers.MapUrlController(app)
	app.Logger.Fatal(app.Start(url))
}