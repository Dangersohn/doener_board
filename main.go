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
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	router.GET("/", view)
	router.GET("/api", api)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func view(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl = template.Must(template.ParseGlob("template/*.html"))
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println(r.Form["sosse1"])
	fmt.Println(r.Form["sosse2"])
}
