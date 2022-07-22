package routes

import (
	"github.com/backsoul/groot/configs"
	"github.com/backsoul/groot/internal/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", controllers.ControllerRegister)
	app.Post("/login", controllers.ControllerLogin)

	// set variables google providers
	goth.UseProviders(
		google.New(configs.Get("GOOGLE_CLIENT_ID"), configs.Get("GOOGLE_CLIENT_SECRET"), "https://auth.backsoul.xyz/auth/callback", "email", "profile"),
		facebook.New(configs.Get("FACEBOOK_CLIENT_ID"), configs.Get("FACEBOOK_CLIENT_SECRET"), "https://auth.backsoul.xyz/auth/callback", "email"),
	)
	app.Get("/auth", goth_fiber.BeginAuthHandler)
	app.Get("/auth/callback", controllers.ControllerAuthCallback)
	// static file html
	app.Static("/", "./public")
}
