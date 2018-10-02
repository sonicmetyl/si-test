package routes

import (
	"github.com/labstack/echo"
	"go-sql/app/handlers"
	"io"
	"html/template"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func SetRoute(e *echo.Echo) *echo.Echo {
	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	e.GET("/", handlers.GetIndex).Name = "index"
	e.GET("/new", handlers.NewTax).Name = "new_tax"
	e.POST("/new", handlers.CreateTax)
	return e
}
