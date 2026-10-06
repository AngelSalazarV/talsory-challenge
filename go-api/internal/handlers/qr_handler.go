package handlers

import (
	"github.com/AngelSalazarV/talsory-challenge/go-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func CalculateQR(c *fiber.Ctx) error {
	var request models.MatrixRequest

	if err:= c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
			"error": "invalid request body",
		})
	} 

	return c.JSON(fiber.Map{
		"matrix": request.Matrix,
	})
}