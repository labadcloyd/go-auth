package database

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-auth/models"
)

var DB *gorm.DB

func Connect() {
	
	connection, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/goauth"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	log.Println("Successfully connected to database")

	DB = connection

	connection.AutoMigrate(&models.User{})
}