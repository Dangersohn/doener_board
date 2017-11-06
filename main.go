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

type Doener struct {
	Kuerzel string
	Sosse1  string
	Sosse2  string
	Sosse3  string
	Salat1  string
	Salat2  string
	Salat3  string
	Salat4  string
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
	e.Static("/images/*", "images")
	log.Fatal(e.Start(":8000"))
}

func show(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}


func api(c echo.Context) error {
	doener := Doener{
		Kuerzel: c.QueryParam("kuerzel"),
		Sosse1:  c.QueryParam("sosse1"),
		Sosse2:  c.QueryParam("sosse2"),
		Sosse3:  c.QueryParam("sosse3"),
		Salat1:  c.QueryParam("salat1"),
		Salat2:  c.QueryParam("salat2"),
		Salat3:  c.QueryParam("salat3"),
		Salat4:  c.QueryParam("salat4"),
	}
	return c.Render(http.StatusOK, "doener.html", doener)
}

//func api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	r.ParseForm()
//	fmt.Println(r.Form["salat"])
//	fmt.Println(r.Form)
//}
