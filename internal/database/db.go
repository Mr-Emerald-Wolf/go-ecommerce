package database

import (
	"fmt"

	"github.com/mr-emerald-wolf/go-api-fiber/models"
	"github.com/mr-emerald-wolf/go-ecommerce/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitializeDB initializes the database connection.
func InitializeDB(cfg *config.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)

	// Connect to the database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Automatically create tables based on the defined models
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	// Store the database connection for later use
	db = database

	return database, nil
}

// GetDB returns the initialized database connection.
func GetDB() *gorm.DB {
	return db
}
