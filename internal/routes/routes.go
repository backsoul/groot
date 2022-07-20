package routes

import (
	"github.com/backsoul/groot/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.ControllerRegister)
}
