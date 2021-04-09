package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")

	switch r.URL.Path {
	case "/":
		fmt.Fprint(rw, "Welcome to my gallery")
	case "/about":
		fmt.Fprint(rw, "About page")
	default:
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw, "Page not found")
	}

}

func main() {
	http.HandleFunc("/", handleFunc)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Panic(err)
	}
}
