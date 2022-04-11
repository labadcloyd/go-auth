package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-auth/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return controllers.Hello(c)
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})
}