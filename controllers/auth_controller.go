package controllers

import (
	"gin-app/db"     // Import your db package for the database connection
	"gin-app/models" // Adjust the import path according to your project structure
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Secret key used for signing tokens
var jwtSecret = []byte("your_secret_key") // Change this to a strong secret key

// LoginRequest defines the structure of the request payload for logging in
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AuthController handles authentication-related requests
type AuthController struct{}

// NewAuthController creates a new instance of AuthController
func NewAuthController() *AuthController {
	return &AuthController{}
}

// GenerateToken creates a new JWT token for the authenticated user
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expiration time (72 hours)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Login handles user login requests
func (ac *AuthController) Login(c *gin.Context) {
	var loginReq LoginRequest

	// Bind the JSON request to the LoginRequest struct
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		log.Printf("Invalid input: %s", err.Error()) // Log the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("Login attempt for username: %s", loginReq.Username) // Log the username

	// Find the user by username using the db package's DB connection
	var user models.Users
	if err := db.DB.Where("username = ?", loginReq.Username).First(&user).Error; err != nil {
		log.Printf("User not found or error occurred: %s", err.Error()) // Log error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Directly check the password without hashing
	if user.Password != loginReq.Password {
		log.Printf("Invalid password attempt for username: %s", loginReq.Username) // Log invalid password
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Log the successful login
	if err := LogUserLogin(user.UserID); err != nil { // Function to log the user login attempt
		log.Printf("Failed to log login attempt for user ID %d: %s", user.UserID, err.Error()) // Log error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log login attempt"})
		return
	}

	// Generate a token
	token, err := GenerateToken(user.UserID)
	if err != nil {
		log.Printf("Failed to generate token for user ID %d: %s", user.UserID, err.Error()) // Log token generation error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Respond with the token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token, // The generated JWT token is returned here
	})
}

// LogUserLogin records a new login event for an existing user
func LogUserLogin(uid uint) error {
	userLogin := models.UserLogin{
		UserID:      uid,
		ModifiedTime: time.Now(), // Set the current timestamp
	}

	// Create a new entry in the UserLogin table
	if err := db.DB.Create(&userLogin).Error; err != nil {
		log.Printf("Error logging user login: %s", err.Error()) // Log error
		return err // Handle the error appropriately
	}

	log.Printf("User login recorded for user ID %d at %s", uid, userLogin.ModifiedTime) // Log successful login
	return nil // Successfully logged the login event
}
