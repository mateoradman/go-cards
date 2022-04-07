package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateoradman/go-cards/internal/models"
)

func main() {
	router := mux.NewRouter()

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)
}
