package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"members-api/pkg/db"
	"members-api/pkg/handlers"
)

func main() {
	DB := db.Init()
    h := handlers.New(DB)
    router := mux.NewRouter()

    // Here we'll define our api endpoints
	router.HandleFunc("/members", handlers.GetAllMembers).Methods(http.MethodGet)
	router.HandleFunc("/members", h.AddMember).Methods(http.MethodPost)

    log.Println("API is running!")
    http.ListenAndServe(":4000", router)
}