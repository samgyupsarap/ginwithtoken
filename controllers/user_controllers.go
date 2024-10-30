package controllers

import (
	"gin-app/db" // Import the db package
	"gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    var samples []struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }
    // Fetch users from the sample table
    if err := db.DB.Model(&models.Sample{}).Select("name, age").Find(&samples).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
        return
    }
    c.JSON(http.StatusOK, samples) // Return the users
}
