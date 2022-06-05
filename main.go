package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pluswhale/go-auth-react/database"
	"github.com/pluswhale/go-auth-react/routes"
)

func main() {

	database.Connect()
	app := fiber.New()

	routes.SetUp(app)

	app.Listen(":8001")
}
