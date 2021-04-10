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
var templateDir string = "views"
var extention string = "gohtml"

func NewView(files ...string) *View {
	files = append(appendToFiles(files), getLaouts()...)
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

func appendToFiles(files []string) []string {
	var finalFiles []string
	for _, v := range files {
		finalFiles = append(finalFiles, fmt.Sprintf("%s/%s.%s", templateDir, v, extention))
	}
	return finalFiles
}
