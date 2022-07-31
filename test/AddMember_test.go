package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"members-api/pkg/db"
	"members-api/pkg/handlers"
	"members-api/pkg/models"

	"github.com/google/uuid"
)


func Test_AddMember_Success(t *testing.T) {
	//**
	//** Setup
	//**
	// Data
	member := models.Member{
		Id: uuid.New(),
		LastName: "Case",
		FirstName: "Test",
		Email: "test@testing.com",
		Mailing: "1234 Mailing Dr., Test, TDD, 123245",
		Phone: "1234567890",
	}
	body, err := json.Marshal(&member)
	if err != nil {
		t.Fatal(err)
	}

	//request
	req, err := http.NewRequest("POST", "/member", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	//Recorder to satisfy the response writer
	rr := httptest.NewRecorder();

	DB := db.Init()
    h := handlers.New(DB)
	handler := http.HandlerFunc(h.AddMember)

	//**
	//** Execution
	//**
	handler.ServeHTTP(rr, req)

	//**
	//** Validation
	//**
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}