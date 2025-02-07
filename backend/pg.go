package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

type Match struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Score string `json:"score"`
}

func main() {
    dsn := "host=localhost user=youruser password=yourpassword dbname=tennisdb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto-migrate the schema
    db.AutoMigrate(&Match{})

    log.Println("Connected to database!")
}
