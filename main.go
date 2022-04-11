package main

import (
	"go-auth/database"
	"go-auth/routes"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
    database.Connect()
    app := fiber.New()

    routes.Setup(app)

    // returning 404 after wrong route
	app.Use(func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusNotFound).SendString("Error 404: not found")
    })

    log.Fatal(app.Listen(":3000"))
}