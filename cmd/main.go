package main

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	"github.com/vnkot/json-kitty/pkg/jsonkitty"
	"github.com/vnkot/json-kitty/pkg/middleware"
)

var staticPath = "static"
var templates = template.Must(template.ParseFiles("templates/index.html"))

func getNewTextAreaResult(value string) string {
	safeValue := html.EscapeString(value)
	return fmt.Sprintf(` 
		<textarea id="json-editor" placeholder="Введите ваш json здесь" name="client-json">%s</textarea>`, safeValue)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func jsonExampleHandler(w http.ResponseWriter, r *http.Request) {
	randomJsonExample := jsonkitty.Examples[rand.Intn(len(jsonkitty.Examples))]

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(getNewTextAreaResult(randomJsonExample)))
}

func jsonFormatHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientJson := r.FormValue("client-json")
	prettyJson, err := jsonkitty.Pretty(clientJson)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(prettyJson)
}

func main() {
	http.Handle("/", middleware.CacheControl(http.HandlerFunc(indexHandler)))
	http.Handle(fmt.Sprintf("/%s/", staticPath), middleware.CacheControl(http.StripPrefix(fmt.Sprintf("/%s/", staticPath), http.FileServer(http.Dir(staticPath)))))

	http.HandleFunc("POST /api/json-format", jsonFormatHandler)
	http.HandleFunc("GET /api/json-example", jsonExampleHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
