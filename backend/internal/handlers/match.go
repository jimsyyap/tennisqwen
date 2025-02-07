package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "tennisapp/backend/internal/models"
    "tennisapp/internal/services"
)

func GetMatchByID(c *gin.Context) {
    id := c.Param("id")
    matchID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID"})
        return
    }

    match, err := services.GetMatchByID(matchID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Match not found"})
        return
    }

    c.JSON(http.StatusOK, match)
}
