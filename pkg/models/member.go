package models

type Member struct {
	Id			string	`json:"id" gorm:"primaryKey"` 
	LastName	string `json:"lastName"`
	FirstName	string `json:"firstName"`
	Email 		string `json:"email"`
	Mailing		string `json:"mailing"`
	Phone		string `json:"phone"`
}