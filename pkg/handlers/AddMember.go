package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"

	"members-api/pkg/models"
)

func (h handler) AddMember(w http.ResponseWriter, r *http.Request) {
    // Read to request body
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
		log.Fatalln(err)
    }
	
    var member models.Member
    json.Unmarshal(body, &member)
	
	//Creatr UUID
	Id := uuid.New()
	member.Id = Id;
	
    //Append to the Members table
    if result := h.DB.Create(&member); result.Error != nil {
        fmt.Println(result.Error)
    }

    // Send a 201 created response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}