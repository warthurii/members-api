package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"members-api/pkg/db"
	"members-api/pkg/handlers"
)

func main() {
	DB := db.Init();
    h := handlers.New(DB);
    router := mux.NewRouter();

    //***************
	//** Endpoints **
	//***************
	router.HandleFunc("/members", h.GetAllMembers).Methods(http.MethodGet);
	router.HandleFunc("/member", h.AddMember).Methods(http.MethodPost);
	router.HandleFunc("/member/", h.DeleteMemberByName).Methods(http.MethodDelete).Queries("lastName", "{lastName}", "firstName", "{firstName}");
	router.HandleFunc("/member/", h.DeleteMemberById).Methods(http.MethodDelete).Queries("id", "{id}");
	router.HandleFunc("/member/", h.GetMemberByName).Methods(http.MethodGet).Queries("lastName", "{lastName}", "firstName", "{firstName}");
	router.HandleFunc("/member/", h.GetMemberById).Methods(http.MethodGet).Queries("id", "{id}");

    log.Println("API is running!");
    http.ListenAndServe(":4000", router);
}