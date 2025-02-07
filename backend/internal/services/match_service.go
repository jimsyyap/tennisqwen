package services

import (
    "tennis-match-analysis/backend/internal/models"
    "tennis-match-analysis/backend/pkg/db"
)

func GetMatchByID(id int) (*models.Match, error) {
    var match models.Match
    result := db.DB.First(&match, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &match, nil
}
