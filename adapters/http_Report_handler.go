package adapters

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
)

type HttpReportHandler struct {
	ReportUsecase usecases.ReportUseCase
}

func NewHttpReportHandler(useCase usecases.ReportUseCase) *HttpReportHandler {
	return &HttpReportHandler{ReportUsecase: useCase}
}

func (h *HttpReportHandler) CreateReport(c *fiber.Ctx) error {
	var reports []entity.Report
	if err := c.BodyParser(&reports); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid report"})
	}

	// Validate each report
	for i, _ := range reports {

		reports[i].ReportDate = time.Now()
		reports[i].UpdatedAt = time.Now()
	}
	if err := h.ReportUsecase.CreateReport(reports); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "reports created successfully"})
}

func (h *HttpReportHandler) GetReports(c *fiber.Ctx) error {
	dateStr := c.Query("date")
	if dateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "date is required"})
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid date format"})
	}

	day := date.Day()
	month := int(date.Month())
	year := date.Year()

	reports, err := h.ReportUsecase.GetReports(day, month, year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(reports)
}

func (h *HttpReportHandler) GetReports7Days(c *fiber.Ctx) error {
	reports, err := h.ReportUsecase.GetReports7Days()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(reports)
}

func (h *HttpReportHandler) GetReports1Month(c *fiber.Ctx) error {
	reports, err := h.ReportUsecase.GetReports1Month()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(reports)
}

func (h *HttpReportHandler) GetReports1Year(c *fiber.Ctx) error {
	reports, err := h.ReportUsecase.GetReports1Year()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(reports)
}
