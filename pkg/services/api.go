package services

import (
	"github.com/backsoul/groot/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func InitializeApi() {
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
