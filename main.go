package main

import (
	"fmt"
	"gogal/controllers"
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

func main() {
	userController := controllers.NewUser()
	pagesController := controllers.NewPage()

	r := mux.NewRouter()

	r.HandleFunc("/", pagesController.Index)
	r.HandleFunc("/about", pagesController.About)
	r.HandleFunc("/register", userController.New).Methods("GET")
	r.HandleFunc("/register", userController.Create).Methods("POST")
	r.HandleFunc("/login", userController.Login)

	r.NotFoundHandler = &_404{}

	hasError(http.ListenAndServe(":8989", r))

}

func hasError(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
