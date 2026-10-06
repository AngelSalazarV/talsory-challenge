package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var request LoginRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if request.Username != "admin" || request.Password != "admin123" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": request.Username,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	signedToken, err := token.SignedString([]byte("secret-key"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": signedToken,
	})
}