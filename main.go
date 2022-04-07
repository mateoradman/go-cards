package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)
}
