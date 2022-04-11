package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Get("/:value", func(c *fiber.Ctx) error {
        return c.SendString("value: " + c.Params("value"))
    })

    log.Fatal(app.Listen(":3000"))
}