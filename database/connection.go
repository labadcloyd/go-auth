package database

import (
	"go-auth/helpers"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-auth/models"
)

var DB *gorm.DB

func Connect() {
	dbUrl := helpers.GoDotEnvVariable("MYSQL_URL")
	
	connection, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	log.Println("Successfully connected to database")

	DB = connection

	connection.AutoMigrate(&models.User{})
}