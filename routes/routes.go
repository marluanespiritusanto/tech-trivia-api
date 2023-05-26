package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marluanespiritusanto/tech-trivia-api/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/trivias", handlers.ListTrivias)
	app.Post("/trivias", handlers.CreateTrivia)
}
