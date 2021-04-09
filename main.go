package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type _404 struct {
}

func (notFound *_404) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(http.StatusNotFound)

	fmt.Fprint(rw, "Sorry, the page you are requesting does not exist!")
}

func home(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprint(rw, "Welcome to my gallery")

}

func about(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprint(rw, "About my gallery")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)

	r.NotFoundHandler = &_404{}

	err := http.ListenAndServe(":8989", r)
	if err != nil {
		log.Panic(err)
	}
}
