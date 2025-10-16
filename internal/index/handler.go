package index

import (
	"fmt"
	"html"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"

	htmlelementstate "github.com/vnkot/json-kitty/pkg/html_element_state"
	"github.com/vnkot/json-kitty/pkg/jsonkitty"
)

type Index struct {
	defaultPageState IndexPageState
}

func NewHandler() *Index {
	return &Index{
		defaultPageState: IndexPageState{
			JsonEditorState{
				FormatButton: htmlelementstate.ButtonState{
					Disabled: true,
				},
			},
			JsonResultState{
				CopyButton: htmlelementstate.ButtonState{
					Disabled: true,
				},
			},
		},
	}
}

func (h *Index) Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", h.defaultPageState)
}

func (h *Index) JSONFormat(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientJson := r.FormValue("json")
	prettyJson, err := jsonkitty.Pretty(clientJson)
	quotedJson := html.EscapeString(strconv.Quote(string(prettyJson)))

	if err != nil {
		templates.ExecuteTemplate(w, "json_result.html", JsonResultState{
			CopyButton: htmlelementstate.ButtonState{
				Disabled: true,
			},
			JsonResult: htmlelementstate.ElementState{
				Children: err.Error(),
			},
		})
		return
	}

	templates.ExecuteTemplate(w, "json_result.html", JsonResultState{
		CopyButton: htmlelementstate.ButtonState{
			Disabled: false,
			OnClick:  template.JS(fmt.Sprintf("navigator.clipboard.writeText(%s).then(() => alert('Скопировано!')).catch(err => alert('Что-то сломалось'))", quotedJson)),
		},
		JsonResult: htmlelementstate.ElementState{
			Children: string(prettyJson),
		},
	})
}

func (h *Index) JSONExample(w http.ResponseWriter, r *http.Request) {
	randomJsonExample := jsonkitty.Examples[rand.Intn(len(jsonkitty.Examples))]

	templates.ExecuteTemplate(w, "json_editor.html", JsonEditorState{
		FormatButton: htmlelementstate.ButtonState{
			Disabled: false,
		},
		JsonTextArea: htmlelementstate.ElementState{
			Children: randomJsonExample,
		},
	})
}
