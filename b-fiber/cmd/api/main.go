package main

import (
	"log"

	"github.com/antoniougarte/b-fiber/internal/handler"
	"github.com/antoniougarte/b-fiber/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,http://localhost:5174,http://localhost:4000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	// Rutas públicas
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Fiber API funcionando correctamente",
		})
	})

	app.Post("/api/auth/login", handler.Login)

	// Rutas protegidas con JWT
	protected := app.Group("/api", middleware.JWTProtected())
	protected.Post("/matrix/rotate", handler.RotateMatrix)

	log.Println("Servidor ejecutándose en http://localhost:4000")
	app.Listen(":4000")
}
