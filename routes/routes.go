package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pluswhale/go-auth-react/controller"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
}
