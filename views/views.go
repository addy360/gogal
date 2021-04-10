package views

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var layout string = "app"
var layoutDir string = "views/layouts"

func NewView(files ...string) *View {

	files = append(files, getLaouts()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func getLaouts() []string {
	files, err := filepath.Glob(fmt.Sprintf("%s/*.gohtml", layoutDir))
	if err != nil {
		log.Panic(err)
	}
	return files
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(rw http.ResponseWriter, data interface{}) error {
	rw.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(rw, v.Layout, data)
}
