package main 

import (
	"github.com/AngelSalazarV/talsory-challenge/go-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app:= fiber.New()

	app.Get("/health", 	func(c *fiber.Ctx) error  {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/api/qr", handlers.CalculateQR)

	app.Listen(":3000")
}