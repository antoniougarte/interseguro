package handler

import (
	"github.com/antoniougarte/b-fiber/internal/model"
	"github.com/antoniougarte/b-fiber/internal/service"
	"github.com/antoniougarte/b-fiber/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "invalid request body")
	}

	// Validar credenciales
	if !service.ValidateCredentials(req.Username, req.Password) {
		return utils.SendError(c, fiber.StatusUnauthorized, "invalid credentials")
	}

	// Generar token
	token, err := service.GenerateJWT(req.Username)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "error generating token")
	}

	return utils.SendSuccess(c, model.LoginResponse{
		Token: token,
		User:  req.Username,
	})
}
