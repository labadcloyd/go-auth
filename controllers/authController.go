package controllers

import (
	"go-auth/models"
	"go-auth/database"
	"log"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data = map[string]string{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	user := models.User {
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}

	log.Println("Successfully registered user")
	return c.JSON(user)
}