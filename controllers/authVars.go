package controllers

import "go-auth/helpers"

var SecretKey = helpers.GoDotEnvVariable("SECRET_KEY")

type ReqSignUp struct {
	Name        string  `json:"name" validate:"required,min=3,max=32"`
	Email       string  `json:"email" validate:"required,email,min=6,max=32"`
	Password		string 	`json:"password" validate:"required,min=10,max=32,requiredCharacters"`
}
type ReqLogin struct {
	Email       string  `json:"email" validate:"required,email,min=6,max=32"`
	Password		string 	`json:"password" validate:"required,min=10,max=32,requiredCharacters"`
}
