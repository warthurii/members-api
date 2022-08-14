package handlers

import (
	"encoding/json"
	"fmt"
	"members-api/pkg/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h handler) GetMemberById(w http.ResponseWriter, r *http.Request) {
    //Get names from request url
    params := mux.Vars(r);
    idString := params["id"];
    
	id, _ := uuid.Parse(idString);
	
    var member models.Member;
    member.Id = id;

    //Find based on id
    if result := h.DB.First(&member, member); result.Error != nil {
        fmt.Println(result.Error);
    }

    //Send Response
    w.Header().Add("Content-Type", "application/json");
    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(member);
}