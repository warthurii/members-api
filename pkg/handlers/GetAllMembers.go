package handlers

import (
	"encoding/json"
	"net/http"

	"members-api/pkg/mocks"
)

func GetAllMembers(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(mocks.Members)
}