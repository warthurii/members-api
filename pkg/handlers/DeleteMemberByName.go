package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"members-api/pkg/models"
)

func (h handler) DeleteMemberByName(w http.ResponseWriter, r *http.Request) {
    //Get names from request url
    params := mux.Vars(r)
    first := params["firstName"]
    lst := params["lastName"]
    
    var member models.Member
    member.FirstName = first
    member.LastName = lst

    //Delete based on name
    if result := h.DB.First(&member, member); result.Error != nil {
        fmt.Println(result.Error)
    }

    h.DB.Delete(&member)
    //Send Response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("deleted")
}