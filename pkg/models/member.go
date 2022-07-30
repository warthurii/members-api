package models

import "github.com/google/uuid"

type Member struct {
	Id			uuid.UUID	`json:"id" gorm:"primaryKey"` 
	LastName	string `json:"lastName"`
	FirstName	string `json:"firstName"`
	Email 		string `json:"email"`
	Mailing		string `json:"mailing"`
	Phone		string `json:"phone"`
}