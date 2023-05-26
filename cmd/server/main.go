package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/marluanespiritusanto/tech-trivia-api/database"
	"github.com/marluanespiritusanto/tech-trivia-api/routes"
)

var APP_PORT = os.Getenv("APP_PORT")

func main() {
	app := fiber.New()

	database.DbConnect()
	routes.SetupRoutes(app)

	app.Listen(":" + APP_PORT)
}
