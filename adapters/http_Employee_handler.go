package adapters

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
)

type HttpEmployeeHandler struct {
	EmployeeUseCase usecases.EmployeeUseCase
}

func NewHttpEmployeeHeadler(useCase usecases.EmployeeUseCase) *HttpEmployeeHandler {
	return &HttpEmployeeHandler{EmployeeUseCase: useCase}
}

func (h *HttpEmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	var Employee entity.Employee
	if err := c.BodyParser(&Employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := h.EmployeeUseCase.CreateEmployee(Employee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "success"})
}

func (h *HttpEmployeeHandler) FindEmployee(c *fiber.Ctx) error {
	name := c.Params("name")
	employee, err := h.EmployeeUseCase.FindEmployee(name)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}
	return c.JSON(employee)
}

func (h *HttpEmployeeHandler) GetEmployees(c *fiber.Ctx) error {
	employee, err := h.EmployeeUseCase.GetEmployees()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}
	return c.JSON(employee)
}

func (h *HttpEmployeeHandler) UpdateUseCaseEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	var Employee entity.Employee
	if err := c.BodyParser(&Employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	Employee.EmployeeID = uint(id)
	if err := h.EmployeeUseCase.UpdateUseCaseEmployee(Employee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success"})
}

func (h *HttpEmployeeHandler) DeleteUseCaseEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}
	if err := h.EmployeeUseCase.DeleteUseCaseEmployee(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success"})
}
