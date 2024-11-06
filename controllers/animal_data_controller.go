// controllers/animal_data_controller.go
package controllers

import (
    "gin-app/db"       // Adjust the import path according to your project structure
    "gin-app/models"   // Import your models package for the Animal struct
    "net/http"

    "github.com/gin-gonic/gin"
)

// GetAnimals retrieves all animals without their IDs
func GetAnimals(c *gin.Context) {
    var animals []models.Animals

    // Use Select to exclude the ID column
    if err := db.DB.Select("animal_name", "age", "color").Find(&animals).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve animals"})
        return
    }

    c.JSON(http.StatusOK, animals)
}
