package adapters

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
)

type HttpDiscountHandler struct {
	DiscountUseCases usecases.DiscountUseCase
}

func NewHttpDiscountHandler(useCases usecases.DiscountUseCase) *HttpDiscountHandler {
	return &HttpDiscountHandler{DiscountUseCases: useCases}
}

func (h *HttpDiscountHandler) CreateDiscount(c *fiber.Ctx) error {
	var Discount entity.Discount
	if err := c.BodyParser(&Discount); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := h.DiscountUseCases.CreateDiscount(Discount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "success"})
}

func (h *HttpDiscountHandler) FindDiscountByCode(c *fiber.Ctx) error {
	Codes := c.Params("code")
	Discount, err := h.DiscountUseCases.FindDiscountByCode(Codes)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Discount not found"})
	}
	return c.JSON(Discount)
}

func (h *HttpDiscountHandler) GetDiscounts(c *fiber.Ctx) error {
	Discount, err := h.DiscountUseCases.GetDiscounts()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Discount not found"})
	}
	return c.JSON(Discount)
}

func (h *HttpDiscountHandler) UpdateUseCaseDiscount(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	var Discount entity.Discount
	if err := c.BodyParser(&Discount); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	Discount.DiscountID = uint(id)
	if err := h.DiscountUseCases.UpdateUseCaseDiscount(Discount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success"})
}

func (h *HttpDiscountHandler) DeleteUseCaseDiscount(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}
	if err := h.DiscountUseCases.DeleteUseCaseDiscount(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success"})
}
