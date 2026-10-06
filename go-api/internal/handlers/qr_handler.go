package handlers

import (
	"github.com/AngelSalazarV/talsory-challenge/go-api/internal/models"
	"github.com/AngelSalazarV/talsory-challenge/go-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

func CalculateQR(c *fiber.Ctx) error {
	var request models.MatrixRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	result, err := services.CalculateQR(request.Matrix)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	stats, err := services.SendMatrixToNode(request.Matrix)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "Node API unavailable",
		})
	}

	return c.JSON(fiber.Map{
		"q":     result.Q,
		"r":     result.R,
		"stats": stats,
	})
}