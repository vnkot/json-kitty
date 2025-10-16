package htmlelementstate

import "html/template"

type ButtonState struct {
	Disabled bool
	OnClick  template.JS
}
