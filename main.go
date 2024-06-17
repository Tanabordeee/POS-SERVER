package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/tanabordeee/pos/adapters"
	"github.com/tanabordeee/pos/entity"
	"github.com/tanabordeee/pos/usecases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func checkMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claim := user.Claims.(jwt.MapClaims)

	if claim["auth_name"] != os.Getenv("ADMINNAME") {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), dbPort, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&entity.Employee{}, &entity.Report{}, &entity.Product{}, &entity.Discount{}, &entity.Auth{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	//Auth
	AuthRepo := adapters.NewGormAuthRepository(db)
	AuthService := usecases.NewAuthService(AuthRepo)
	AuthHandler := adapters.NewHttpAuthHandler(AuthService)

	// EMPLOYEE
	EmployeeRepo := adapters.NewGormEmployeeRepository(db)
	EmployeeService := usecases.NewEmployeeService(EmployeeRepo)
	EmployeeHandler := adapters.NewHttpEmployeeHeadler(EmployeeService)

	app.Post("/Createadmin", EmployeeHandler.CreateEmployee)
	app.Post("/CreateAuth", AuthHandler.CreateAuth)
	app.Post("/GetAuth", AuthHandler.CheckAuth)

	// jwtCheck
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWTSECRETKEY")),
	}))

	// GROUP MIDDLEWARE
	AdminGroup := app.Group("/AdminGroup")
	AdminGroup.Use(checkMiddleware)

	AdminGroup.Post("/CreateAuth", AuthHandler.CreateAuth)
	AdminGroup.Put("/UpdateAuth/:id", AuthHandler.UpdateAuth)
	AdminGroup.Delete("/DeleteAuth/:id", AuthHandler.DeleteAuth)

	AdminGroup.Post("/CreateEmployee", EmployeeHandler.CreateEmployee)
	AdminGroup.Get("/FindEmployee/:name", EmployeeHandler.FindEmployee)
	AdminGroup.Get("/GetEmployee", EmployeeHandler.GetEmployees)
	AdminGroup.Put("/UpdateEmployee/:id", EmployeeHandler.UpdateUseCaseEmployee)
	AdminGroup.Delete("/DeleteEmployee/:id", EmployeeHandler.DeleteUseCaseEmployee)

	//REPORT
	ReportRepo := adapters.NewGormReportRepository(db)
	ReportService := usecases.NewReportService(ReportRepo)
	ReportHandler := adapters.NewHttpReportHandler(ReportService)

	app.Post("/CreateReport", ReportHandler.CreateReport)
	AdminGroup.Get("/GetReport", ReportHandler.GetReports)
	AdminGroup.Get("/GetReports7Days", ReportHandler.GetReports7Days)
	AdminGroup.Get("/GetReports1Month", ReportHandler.GetReports1Month)
	AdminGroup.Get("/GetReports1Year", ReportHandler.GetReports1Year)

	//PRODUCT
	ProductRepo := adapters.NewGormProductRepository(db)
	ProductService := usecases.NewProductService(ProductRepo)
	ProducHandler := adapters.NewHttpProductHandler(ProductService)

	AdminGroup.Post("/CreateProduct", ProducHandler.CreateProduct)
	app.Get("/GetProduct", ProducHandler.GetProducts)
	app.Get("/FindProductByName/:name", ProducHandler.FindProductByName)
	AdminGroup.Put("/UpdateProduct/:id", ProducHandler.UpdateUseCaseProduct)
	AdminGroup.Delete("/DeleteProduct/:id", ProducHandler.DeleteUseCaseProduct)

	//DISCOUNT
	DiscountRepo := adapters.NewGormDiscountRepository(db)
	DiscountService := usecases.NewDiscountService(DiscountRepo)
	DiscountHandler := adapters.NewHttpDiscountHandler(DiscountService)

	AdminGroup.Post("/CreateDiscount", DiscountHandler.CreateDiscount)
	app.Get("/FindDiscountByCode/:code", DiscountHandler.FindDiscountByCode)
	app.Get("/GetDiscounts", DiscountHandler.GetDiscounts)
	AdminGroup.Put("/UpdateUseCaseDiscount/:id", DiscountHandler.UpdateUseCaseDiscount)
	AdminGroup.Delete("/DeleteUseCaseDiscount/:id", DiscountHandler.DeleteUseCaseDiscount)

	// Start server in a goroutine
	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// Wait for the server to start
	time.Sleep(2 * time.Second)

	// Run initialization to call the APIs once
	initializeAPIs(db)

	// Keep the main goroutine alive
	select {}
}

func initializeAPIs(db *gorm.DB) {
	// Check if admin already exists
	var count int64
	db.Model(&entity.Employee{}).Where("role = ?", "admin").Count(&count)
	if count > 0 {
		log.Printf("Admin already exists, skipping creation.")
		return
	}

	// Data to be sent in the requests
	adminData := map[string]interface{}{
		"name":   "admin",
		"role":   "admin",
		"salary": 9999,
	}
	authData := map[string]interface{}{
		"employee_id": 1,
		"username":    "admin",
		"password":    "admin123",
	}

	// Send Createadmin request
	if err := sendRequest("/Createadmin", adminData); err != nil {
		log.Printf("Skipping /Createadmin due to error: %v", err)
	}

	// Send CreateAuth request
	if err := sendRequest("/CreateAuth", authData); err != nil {
		log.Printf("Skipping /CreateAuth due to error: %v", err)
	}
}

// Helper function to send HTTP POST requests
func sendRequest(endpoint string, data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal data: %v", err)
		return nil // Return nil to continue execution
	}

	resp, err := http.Post("http://localhost:8000"+endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to send request to %s: %v", endpoint, err)
		return nil // Return nil to continue execution
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Non-OK HTTP status: %s", resp.Status)
		return nil // Return nil to continue execution
	}

	log.Printf("Successfully called %s", endpoint)
	return nil // Return nil to indicate success
}
