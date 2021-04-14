package main

import (
	"fmt"
	"gogal/controllers"
	"gogal/helpers"
	"gogal/services"
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
	connectionString := "user=postgres dbname=gogal port=5432 sslmode=disable"
	db := helpers.DbConnection(connectionString)
	us := services.NewUserService(db)
	gs := services.NewGalleryService(db)
	userController := controllers.NewUser(us)
	pagesController := controllers.NewPage()
	galleryController := controllers.NewGarrely(gs)
	us.TableRefresh()
	gs.TableRefresh()
	r := mux.NewRouter()

	r.HandleFunc("/", pagesController.Index)
	r.HandleFunc("/about", pagesController.About)
	r.HandleFunc("/register", userController.New).Methods("GET")
	r.HandleFunc("/register", userController.Create).Methods("POST")
	r.HandleFunc("/login", userController.Login).Methods("GET")
	r.HandleFunc("/login", userController.SignIn).Methods("POST")
	r.HandleFunc("/cookie", userController.CookieTest)

	r.HandleFunc("/gallery/show", galleryController.Show)
	r.HandleFunc("/gallery/create", galleryController.Create).Methods("GET")
	r.HandleFunc("/gallery/store", galleryController.CreateGallery).Methods("POST")

	r.NotFoundHandler = &_404{}

	hasError(http.ListenAndServe(":8989", r))

}

func hasError(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
