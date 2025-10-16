package index

import "text/template"

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/subtemplates/json_editor.html", "templates/subtemplates/json_result.html"))
