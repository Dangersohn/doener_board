package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", show)
	e.GET("/api", api)
	log.Fatal(e.Start(":8000"))
}

func show(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func api(c echo.Context) error {
	sosse := c.QueryParam("sosse")
	fmt.Println(sosse)
	return c.String(http.StatusOK, "Soßen: "+sosse)
}

//func api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	r.ParseForm()
//	fmt.Println(r.Form["salat"])
//	fmt.Println(r.Form)
//}
