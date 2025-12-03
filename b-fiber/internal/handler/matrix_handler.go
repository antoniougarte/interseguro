package handler

import (
	"strings"

	"github.com/antoniougarte/b-fiber/internal/model"
	"github.com/antoniougarte/b-fiber/internal/service"
	"github.com/antoniougarte/b-fiber/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func RotateMatrix(c *fiber.Ctx) error {
	var req model.MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "invalid JSON")
	}

	// Obtener token del header (ya validado por el middleware)
	authHeader := c.Get("Authorization")
	var token string
	if authHeader == "" {
		// No debería pasar si middleware JWTProtected valida, pero por seguridad:
		return utils.SendError(c, fiber.StatusUnauthorized, "missing Authorization header")
	}
	// forma esperada: "Bearer <token>"
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = authHeader[len("Bearer "):]
	} else {
		token = authHeader
	}

	resp, err := service.ProcessMatrix(req, token)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccess(c, resp)
}
