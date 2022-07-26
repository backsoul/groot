package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/backsoul/groot/cmd/groot/docs"
	"github.com/backsoul/groot/internal/database"
	"github.com/backsoul/groot/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.InitializeDb()
	app := fiber.New()
	app.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking:  false,
		DocExpansion: "",
		URL:          "/swagger/doc.json",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers",
		AllowMethods:     "POST,PATCH",
	}))
	routes.SetupRoutes(app)
	app.Listen(":8000")
}
