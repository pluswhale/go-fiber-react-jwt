package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pluswhale/go-auth-react/database"
	"github.com/pluswhale/go-auth-react/routes"
)

func main() {

	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetUp(app)

	app.Listen(":8002")
}
