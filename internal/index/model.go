package index

import (
	htmlelementstate "github.com/vnkot/json-kitty/pkg/html_element_state"
)

type JsonEditorState struct {
	FormatButton htmlelementstate.ButtonState
	JsonTextArea htmlelementstate.ElementState
}

type JsonResultState struct {
	CopyButton htmlelementstate.ButtonState
	JsonResult htmlelementstate.ElementState
}

type IndexPageState struct {
	JsonEditorState
	JsonResultState
}
