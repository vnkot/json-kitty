package main

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	"github.com/vnkot/json-kitty/pkg/jsonkitty"
)

var staticPath = "static"
var templates = template.Must(template.ParseFiles("templates/index.html"))

func getNewTextAreaResult(value string) string {
	safeValue := html.EscapeString(value)
	return fmt.Sprintf(` 
		<textarea id="json-editor" placeholder="Введите ваш json здесь" name="client-json">%s</textarea>`, safeValue)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})

	http.HandleFunc("GET /api/json-example", func(w http.ResponseWriter, r *http.Request) {
		randomJsonExample := jsonkitty.Examples[rand.Intn(len(jsonkitty.Examples))]

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(getNewTextAreaResult(randomJsonExample)))
	})

	http.HandleFunc("POST /api/json-format", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.Handle(fmt.Sprintf("/%s/", staticPath), http.StripPrefix(fmt.Sprintf("/%s/", staticPath), http.FileServer(http.Dir(staticPath))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
