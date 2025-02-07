package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using environment variables.")
    }

    // Example: Load database URL
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL is not set")
    }
}
