package middleware

import (
	"strings"

	"github.com/antoniougarte/b-fiber/internal/service"
	"github.com/antoniougarte/b-fiber/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return utils.SendError(c, fiber.StatusUnauthorized, "missing authorization header")
		}

		// Formato: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.SendError(c, fiber.StatusUnauthorized, "invalid authorization format")
		}

		token := parts[1]

		claims, err := service.ValidateJWT(token)
		if err != nil {
			return utils.SendError(c, fiber.StatusUnauthorized, "invalid or expired token")
		}

		// Guardar usuario en el contexto para usar en handlers
		c.Locals("username", claims.Username)

		return c.Next()
	}
}
