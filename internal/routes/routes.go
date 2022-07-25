package routes

import (
	"github.com/backsoul/groot/configs"
	"github.com/backsoul/groot/internal/controllers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{"status": "Groot API Working!"})
		c.SendStatus(fiber.StatusOK)
		return nil
	})
	api.Get("/sessions/oauth/google", controllers.ControllerAuthGoogleProvider)
	// JWT Middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(configs.Get("JWT_KEY")),
	}))
	api.Get("/me", controllers.Me)
	api.Get("/refresh", controllers.Refresh)
}
