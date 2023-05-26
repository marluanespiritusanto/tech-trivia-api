package handlers

import (
	"github.com/marluanespiritusanto/tech-trivia-api/database"
	"github.com/marluanespiritusanto/tech-trivia-api/models"

	"github.com/gofiber/fiber/v2"
)

func ListTrivias(c *fiber.Ctx) error {
	trivias := []models.Trivia{}
	database.DB.Db.Find(&trivias)

	return c.Status(200).JSON(trivias)
}

func CreateTrivia(c *fiber.Ctx) error {
	trivia := new(models.Trivia)
	if err := c.BodyParser(trivia); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})
	}

	database.DB.Db.Create(&trivia)

	return c.Status(200).JSON(trivia)
}
