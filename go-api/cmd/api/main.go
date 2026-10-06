package main 

import (
	"github.com/AngelSalazarV/talsory-challenge/go-api/internal/handlers"

	"github.com/gofiber/fiber/v2"

	"github.com/AngelSalazarV/talsory-challenge/go-api/internal/middleware"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app:= fiber.New()

	app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:3002",
			AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/health", 	func(c *fiber.Ctx) error  {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/auth/login", handlers.Login)

	app.Post("/api/qr", middleware.JWTProtected, handlers.CalculateQR)

	app.Listen(":3000")
}