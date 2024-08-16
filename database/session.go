package database

import (
    "database/sql"
    "log"
    "os"
    "path/filepath"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func NewSession() (*sql.DB, error) {
    // Get the parent directory
    parentDir := filepath.Dir("..")

    // Load the .env file from the parent directory
    err := godotenv.Load(filepath.Join(parentDir, ".env"))
    if err != nil {
        log.Fatalf("Error loading .env file from parent directory: %v", err)
    }

    // Get the DATABASE_URL from the environment variables
    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL == "" {
        log.Fatalf("DATABASE_URL not set in .env file")
    }
    // Open the database connection
    return sql.Open("postgres", databaseURL)
}