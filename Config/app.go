package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading .env file: %s", err)
	}

	// Connect to the database using environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("Error connecting to the database: %s", err)
	}

	// Set up a database connection pool
	d.DB().SetMaxOpenConns(10) // Adjust as needed
	d.DB().SetMaxIdleConns(5)  // Adjust as needed

	db = d
	return nil
}

func GetDB() *gorm.DB {
	return db
}
