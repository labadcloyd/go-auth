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

    log.Fatal(app.Listen(":3000"))
}