package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var staticPath = "static"
var templates = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})

	http.Handle(fmt.Sprintf("/%s/", staticPath), http.StripPrefix(fmt.Sprintf("/%s/", staticPath), http.FileServer(http.Dir(staticPath))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
