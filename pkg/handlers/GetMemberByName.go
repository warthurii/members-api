package handlers

import (
	"encoding/json"
	"fmt"
	"members-api/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (h handler) GetMemberByName(w http.ResponseWriter, r *http.Request) {
    //Get names from request url
    params := mux.Vars(r);
    first := params["firstName"];
    lst := params["lastName"];
    
    var member models.Member;
    member.FirstName = first;
    member.LastName = lst;

    //Retrieve based on name
    if result := h.DB.First(&member, member); result.Error != nil {
        fmt.Println(result.Error);
    }

    //Send Response
    w.Header().Add("Content-Type", "application/json");
    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(member);
}