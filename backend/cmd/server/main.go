package main

import (
    "log"
    "tennisapp/backend/internal/handlers"
    "tennisapp/backend/pkg/db"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize database connection
    db.InitDB()

    // Create a new Gin router
    r := gin.Default()

    // Define API routes
    r.GET("/matches/:id", handlers.GetMatchByID)
    r.POST("/matches", handlers.CreateMatch)

    // Start the server
    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
