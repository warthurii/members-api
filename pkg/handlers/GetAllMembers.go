package handlers

import (
	"encoding/json"
	"net/http"

	"members-api/pkg/models"
)

func (h handler) GetAllMembers(w http.ResponseWriter, r *http.Request) {
	var members []models.Member;
	h.DB.Find(&members)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(members)
}