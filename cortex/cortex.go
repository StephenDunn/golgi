package cortex

import (
	"fmt"
	"golgi/cortex/view"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func Startup() {
	e := echo.New()
	tmpl := template.New("index")

	var err error
	if tmpl, err = tmpl.Parse(view.Index); err != nil {
		fmt.Println(err)
	}

	e.Renderer = &TemplateRenderer{
		templates: tmpl,
	}

	e.GET("/yahoo", hello)

	e.GET("/dothing/:id", func(c echo.Context) error {
		id := string(c.Param("id"))

		output := id + id

		return c.Render(http.StatusOK, "item-count", output)
	})

	e.Logger.Fatal(e.Start(":8091"))
}

func hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index", view.Index)
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
