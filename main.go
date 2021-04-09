package main

import (
	"fmt"
	"gogal/views"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var homeView *views.View

var aboutView *views.View

var registerView *views.View

var loginView *views.View

type _404 struct {
}

func (notFound *_404) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(http.StatusNotFound)

	fmt.Fprint(rw, "Sorry, the page you are requesting does not exist!")
}

func home(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	hasError(homeView.Render(rw, nil))

}

func about(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	hasError(aboutView.Render(rw, nil))
}

func register(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	hasError(registerView.Render(rw, nil))
}

func login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	hasError(loginView.Render(rw, nil))
}

func main() {
	var err error

	homeView = views.NewView("views/home.gohtml")

	aboutView = views.NewView("views/about.gohtml")

	registerView = views.NewView("views/register.gohtml")

	loginView = views.NewView("views/login.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/register", register)
	r.HandleFunc("/login", login)

	r.NotFoundHandler = &_404{}

	err = http.ListenAndServe(":8989", r)
	if err != nil {
		log.Panic(err)
	}
}

func hasError(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
