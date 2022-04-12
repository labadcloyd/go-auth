package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-auth/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/signup", controllers.Signup)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)

}