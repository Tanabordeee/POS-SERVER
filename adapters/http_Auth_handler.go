package adapters

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
)

type HttpAuthHandler struct {
	AuthUseCases usecases.AuthUseCase
}

func NewHttpAuthHandler(useCases usecases.AuthUseCase) *HttpAuthHandler {
	return &HttpAuthHandler{AuthUseCases: useCases}
}

func (h *HttpAuthHandler) CheckAuth(c *fiber.Ctx) error {
	type AuthRequest struct {
		AuthName     string `json:"auth_name"`
		AuthPassword string `json:"auth_password"`
	}

	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	Auth := entity.Auth{
		Username: req.AuthName,
		Password: req.AuthPassword,
	}
	token, err := h.AuthUseCases.CheckAuth(Auth)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"token": token})
}

func (h *HttpAuthHandler) CreateAuth(c *fiber.Ctx) error {
	var Auth entity.Auth
	if err := c.BodyParser(&Auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	err := h.AuthUseCases.CreateAuth(Auth)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "success"})
}

func (h *HttpAuthHandler) UpdateAuth(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	var Auth entity.Auth
	if err := c.BodyParser(&Auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	Auth.AuthID = uint(id)
	if err := h.AuthUseCases.UpdateAuth(Auth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success"})
}

func (h *HttpAuthHandler) DeleteAuth(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}
	if err := h.AuthUseCases.DeleteAuth(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success"})
}
