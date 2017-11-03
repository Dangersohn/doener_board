package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {

}

func main() {
	router := httprouter.New()
	router.GET("/", view)
	router.GET("/api", api)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func view(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl = template.Must(template.ParseGlob("template/*.html"))
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	fmt.Println(r.Form["salat"])
	fmt.Println(r.Form)
}
