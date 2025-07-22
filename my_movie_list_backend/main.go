package main

import (
	"log"
	
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"my_movie_list_backend/database"
	"my_movie_list_backend/routes"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Continuing with system environment variables...")
	}

	// Connect to the database
	if err := database.Connect(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

    if err := database.Migrate(); err != nil {
        log.Panic("Migration failed:", err)
    }
	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	// Start server
	app.Listen(":3000")
	
}
