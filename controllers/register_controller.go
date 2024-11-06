package controllers

import (
	"gin-app/db"     // Import your db package for the database connection
	"gin-app/models" // Adjust the import path according to your project structure
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterRequest defines the structure of the request payload for registration
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"` // Enforce password length
}

// RegisterController handles registration-related requests
type RegisterController struct{}

// NewRegisterController creates a new instance of RegisterController
func NewRegisterController() *RegisterController {
	return &RegisterController{}
}

// Register handles user registration requests
func (rc *RegisterController) Register(c *gin.Context) {
	var regReq RegisterRequest

	// Bind the JSON request to the RegisterRequest struct
	if err := c.ShouldBindJSON(&regReq); err != nil {
		log.Printf("Invalid input: %s", err.Error()) // Log the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("Registration attempt for username: %s", regReq.Username) // Log the username

	// Check if the username already exists
	var existingUser models.UserRegister
	if err := db.DB.Where("username = ?", regReq.Username).First(&existingUser).Error; err == nil {
		log.Printf("Username %s already exists", regReq.Username) // Log duplicate username
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Create the new user entry
	newUser := models.UserRegister{
		Username:  regReq.Username,
		Password:  regReq.Password, // Store plain text password (not recommended)
		CreatedAt: time.Now(),
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		log.Printf("Error creating user: %s", err.Error()) // Log database error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	log.Printf("User %s created successfully", newUser.Username) // Log successful registration

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
	})
}
