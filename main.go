package main

import (
	"fmt"
	"gogal/controllers"
	"gogal/views"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var homeView *views.View

var aboutView *views.View

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

func main() {
	userController := controllers.NewUser()
	var err error

	homeView = views.NewView("views/home.gohtml")

	aboutView = views.NewView("views/about.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/register", userController.New).Methods("GET")
	r.HandleFunc("/register", userController.Create).Methods("POST")
	r.HandleFunc("/login", userController.Login)

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
