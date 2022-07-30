package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"members-api/pkg/handlers"
)

func main() {
    router := mux.NewRouter()

    // Here we'll define our api endpoints
	router.HandleFunc("/members", handlers.GetAllMembers).Methods(http.MethodGet)

    log.Println("API is running!")
    http.ListenAndServe(":4000", router)
}