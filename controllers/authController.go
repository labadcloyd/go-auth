package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data = map[string]string{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	log.Println("Successfully registered user")

	return c.JSON(data)
}