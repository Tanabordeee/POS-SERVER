package adapters

import (
	"io/ioutil"
	"net/url"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
)

type HttpProductHandler struct {
	ProductUsecase usecases.ProductUseCase
}

func NewHttpProductHandler(useCase usecases.ProductUseCase) *HttpProductHandler {
	return &HttpProductHandler{ProductUsecase: useCase}
}

func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	productName := c.FormValue("product_name")
	priceStr := c.FormValue("price")
	costStr := c.FormValue("cost")

	// แปลงค่าจาก string เป็น float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid price"})
	}

	cost, err := strconv.ParseFloat(costStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid cost"})
	}
	// อ่านไฟล์จาก form
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to read file"})
	}

	// แปลงไฟล์เป็น byte
	fileHandle, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to open file"})
	}
	defer fileHandle.Close()

	// แปลงไฟล์เป็น byte
	fileBytes, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to read file bytes"})
	}

	// สร้าง Product object
	product := entity.Product{
		ProductName: productName,
		Price:       price,
		Image:       fileBytes,
		Cost:        cost,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// บันทึกลงฐานข้อมูล
	if err := h.ProductUsecase.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "product created successfully"})
}

func (h *HttpProductHandler) FindProductByName(c *fiber.Ctx) error {
	name := c.Params("name")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product name"})
	}
	product, err := h.ProductUsecase.FindProductByName(decodedName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

func (h *HttpProductHandler) GetProducts(c *fiber.Ctx) error {
	product, err := h.ProductUsecase.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

func (h *HttpProductHandler) UpdateUseCaseProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	productName := c.FormValue("product_name")
	priceStr := c.FormValue("price")
	costStr := c.FormValue("cost")

	// แปลงค่าจาก string เป็น float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid price"})
	}

	cost, err := strconv.ParseFloat(costStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid cost"})
	}
	// อ่านไฟล์จาก form
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to read file"})
	}

	// แปลงไฟล์เป็น byte
	fileHandle, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to open file"})
	}
	defer fileHandle.Close()

	// แปลงไฟล์เป็น byte
	fileBytes, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to read file bytes"})
	}

	// สร้าง Product object
	Product := entity.Product{
		ProductName: productName,
		Price:       price,
		Image:       fileBytes,
		Cost:        cost,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	Product.ProductID = uint(id)
	if err := h.ProductUsecase.UpdateUseCaseProduct(Product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success"})
}

func (h *HttpProductHandler) DeleteUseCaseProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}
	if err := h.ProductUsecase.DeleteUseCaseProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success"})
}
