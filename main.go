package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Template struct {
	templates *template.Template
}

type Doener struct {
	Zeit    string
	Kuerzel string
	Gericht string
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

//DB STUFF
const (
	//MongoDBHost ist der Hostname
	MongoDBHost = "127.0.0.1:27017"
	//DBName ist der Name der DB
	DBName = "test"
	//CollectionName ist der Name der Collection
	CollectionName = "people"
)

var mgoSession *mgo.Session

func GetMongoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(MongoDBHost)
		if err != nil {
			log.Fatal("Failed to start the Mongo session")
		}
	}
	return mgoSession.Clone()
}

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/jana", jana)
	e.GET("/", show)
	e.GET("/box", box)
	e.GET("/api", api)
	e.GET("/orders", orders)
	e.Static("/images/*", "images")
	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func show(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func box(c echo.Context) error {
	return c.Render(http.StatusOK, "nico.html", nil)
}

func jana(c echo.Context) error {
	return c.Render(http.StatusOK, "jana.html", nil)
}

func api(c echo.Context) error {

	s := GetMongoSession()
	defer s.Close()

	doener := Doener{
		Zeit:    time.Now().Format(time.RFC3339Nano),
		Kuerzel: strings.ToUpper(c.QueryParam("kuerzel")),
		Gericht: c.QueryParam("gericht"),
		Sosse1:  c.QueryParam("sosse1"),
		Sosse2:  c.QueryParam("sosse2"),
		Sosse3:  c.QueryParam("sosse3"),
		Salat1:  c.QueryParam("salat1"),
		Salat2:  c.QueryParam("salat2"),
		Salat3:  c.QueryParam("salat3"),
		Salat4:  c.QueryParam("salat4"),
	}
	//j, _ := json.Marshal(doener)
	msesion := s.DB(DBName).C(CollectionName)
	err := msesion.Insert(doener)
	if err != nil {
		fmt.Print(err)
	}

	return c.Render(http.StatusOK, "doener.html", doener)
}

func orders(c echo.Context) error {

	s := GetMongoSession()
	defer s.Close()

	msesion := s.DB(DBName).C(CollectionName)
	var doener []Doener
	err := msesion.Find(bson.M{}).All(&doener)
	if err != nil {
		fmt.Print(err)
	}

	//iter := db.NewIterator(util.BytesPrefix([]byte(time.Now().Format("2006-01-02"))), nil)
	//for iter.Next() {
	//	var d Doener
	//	json.Unmarshal(iter.Value(), &d)

	//	doener = append(doener, d)

	//}
	//iter.Release()

	return c.Render(http.StatusOK, "orders.html", doener)
}
