package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	
	_, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/mysql"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	log.Println("Successfully connected to database")
}