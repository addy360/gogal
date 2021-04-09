package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprint(rw, "Welcome to my gallery")
}

func main() {
	http.HandleFunc("/", handleFunc)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Panic(err)
	}
}
