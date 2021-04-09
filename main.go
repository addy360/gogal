package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

var aboutTemplate *template.Template

type _404 struct {
}

func (notFound *_404) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(http.StatusNotFound)

	fmt.Fprint(rw, "Sorry, the page you are requesting does not exist!")
}

func home(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(rw, nil); err != nil {
		log.Panic(err)
	}

}

func about(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	if err := aboutTemplate.Execute(rw, nil); err != nil {
		log.Panic(err)
	}
}

func main() {
	var err error

	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	aboutTemplate, err = template.ParseFiles("views/about.gohtml")

	if err != nil {
		log.Panic(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)

	r.NotFoundHandler = &_404{}

	err = http.ListenAndServe(":8989", r)
	if err != nil {
		log.Panic(err)
	}
}
