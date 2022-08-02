package test

import (
	"bytes"
	"members-api/pkg/db"
	"members-api/pkg/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_DeleteMemberByName_Success(t *testing.T){
	//**
	//** Setup
	//**
	// Data
	body := ""
	req, err := http.NewRequest("POST", "/member/?lastName=test&firstName=Will", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder();

	DB := db.Init()
    h := handlers.New(DB)
	handler := http.HandlerFunc(h.DeleteMemberByName)

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