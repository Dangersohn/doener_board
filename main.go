package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
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

var db *leveldb.DB

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}

	db, _ = leveldb.OpenFile("db", nil)

	defer db.Close()

	e := echo.New()
	e.Renderer = t
	e.GET("/", show)
	e.GET("/box", box)
	e.GET("/api", api)
	e.GET("/orders", orders)
	e.Static("/images/*", "images")
	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func show(c echo.Context) error {
	return c.Render(http.StatusOK, "new_index.html", nil)
}

func box(c echo.Context) error {
	return c.Render(http.StatusOK, "nico.html", nil)
}

func api(c echo.Context) error {
	doener := Doener{
		Kuerzel: strings.ToUpper(c.QueryParam("kuerzel")),
		Sosse1:  c.QueryParam("sosse1"),
		Sosse2:  c.QueryParam("sosse2"),
		Sosse3:  c.QueryParam("sosse3"),
		Salat1:  c.QueryParam("salat1"),
		Salat2:  c.QueryParam("salat2"),
		Salat3:  c.QueryParam("salat3"),
		Salat4:  c.QueryParam("salat4"),
	}
	j, _ := json.Marshal(doener)

	t := time.Now().Format(time.RFC3339Nano)

	err := db.Put([]byte(t), j, nil)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render(http.StatusOK, "doener.html", doener)
}

func orders(c echo.Context) error {

	var doener []Doener

	iter := db.NewIterator(util.BytesPrefix([]byte(time.Now().Format("2006-01-02"))), nil)
	for iter.Next() {
		var d Doener
		json.Unmarshal(iter.Value(), &d)

		doener = append(doener, d)

	}
	iter.Release()

	return c.Render(http.StatusOK, "orders.html", doener)
}
