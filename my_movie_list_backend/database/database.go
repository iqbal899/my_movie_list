package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my_movie_list_backend/models"
)

var DB *gorm.DB

// Connect initializes the database connection using environment variables
func Connect() error {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not loaded, falling back to system environment variables")
	}

	// Build DSN string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	fmt.Println("✅ Database connected successfully")
	return nil
}

// Migrate auto-migrates the models to the database
func Migrate() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Movie{},
		&models.WatchList{},
	)
	if err != nil {
		return fmt.Errorf("failed to run auto migration: %w", err)
	}
	fmt.Println("✅ Database migration completed")
	return nil
}
