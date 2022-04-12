package models

type User struct {
	Id				uint		`json:"id"`
	Name			string	`json:"name"`
	Email			string	`json:"email" gorm:"unique"`
	Password	[]byte	`json:"-"` // putting a minus means not returning field
}


