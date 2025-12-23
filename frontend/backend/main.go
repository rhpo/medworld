package main

import (
	"log"
	"medworld-backend/database"
	"medworld-backend/middleware"
	"medworld-backend/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	if err := database.InitDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "MedWorld API v1.0",
		ErrorHandler: customErrorHandler,
	})

	// Global middleware
	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.CORSMiddleware())

	// Setup routes
	routes.SetupRoutes(app)

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Start server
	log.Printf("🚀 Server starting on http://localhost:%s", port)
	log.Printf("📚 API documentation: http://localhost:%s/api/v1/health", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// customErrorHandler handles errors globally
func customErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send custom error response
	return c.Status(code).JSON(fiber.Map{
		"message": err.Error(),
		"error":   true,
	})
}
