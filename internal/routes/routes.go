package routes

import (
	"github.com/backsoul/groot/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{"status": "Groot API Working!"})
		c.SendStatus(fiber.StatusOK)
		return nil
	})
	api.Get("/sessions/oauth/google", controllers.ControllerAuthGoogleProvider)

	api.Post("/me", controllers.Me)
}
