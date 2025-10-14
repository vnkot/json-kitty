package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/vnkot/json-kitty/pkg/jsonkitty"
	"github.com/vnkot/json-kitty/pkg/middleware"
)

var staticPath = "static"
var templates = template.Must(template.ParseFiles("templates/index.html", "templates/subtemplates/json_editor.html", "templates/subtemplates/json_result.html"))

type ButtonState struct {
	Disabled bool
	OnClick  template.JS
}

type ElementState struct {
	Children string
}

type JsonEditorState struct {
	FormatButton ButtonState
	JsonTextArea ElementState
}

type JsonResultSate struct {
	CopyButton ButtonState
	JsonResult ElementState
}

type IndexPageState struct {
	JsonEditorState
	JsonResultSate
}

var indexPageState = IndexPageState{
	JsonEditorState{
		FormatButton: ButtonState{
			Disabled: true,
		},
	},
	JsonResultSate{
		CopyButton: ButtonState{
			Disabled: true,
		},
	},
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", indexPageState)
}

func jsonExampleHandler(w http.ResponseWriter, r *http.Request) {
	randomJsonExample := jsonkitty.Examples[rand.Intn(len(jsonkitty.Examples))]

	templates.ExecuteTemplate(w, "json_editor.html", JsonEditorState{
		FormatButton: ButtonState{
			Disabled: false,
		},
		JsonTextArea: ElementState{
			Children: randomJsonExample,
		},
	})
}

func jsonFormatHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientJson := r.FormValue("json")
	prettyJson, err := jsonkitty.Pretty(clientJson)
	quotedJson := strconv.Quote(string(prettyJson))

	if err != nil {
		templates.ExecuteTemplate(w, "json_result.html", JsonResultSate{
			CopyButton: ButtonState{
				Disabled: true,
			},
			JsonResult: ElementState{
				Children: err.Error(),
			},
		})
		return
	}

	templates.ExecuteTemplate(w, "json_result.html", JsonResultSate{
		CopyButton: ButtonState{
			Disabled: false,
			OnClick:  template.JS(fmt.Sprintf("navigator.clipboard.writeText(%s).then(() => alert('Copied!')).catch(err => alert('Error: ' + err))", quotedJson)),
		},
		JsonResult: ElementState{
			Children: string(prettyJson),
		},
	})
}

func main() {
	http.Handle("/", middleware.CacheControl(http.HandlerFunc(indexHandler)))
	http.Handle(fmt.Sprintf("/%s/", staticPath), middleware.CacheControl(http.StripPrefix(fmt.Sprintf("/%s/", staticPath), http.FileServer(http.Dir(staticPath)))))

	http.HandleFunc("POST /api/json-format", jsonFormatHandler)
	http.HandleFunc("GET /api/json-example", jsonExampleHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
