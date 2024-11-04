package main

import (
	"gin-app/controllers" // Adjust the import path according to your project structure
	"gin-app/db"          // Import your db package for the database connection

	"github.com/gin-gonic/gin"
)

func main() {
    // Connect to the database
    db.ConnectDB() // Ensure your db package has a Connect function to initialize the DB connection

    // Initialize Gin
    r := gin.Default()

    // Create an instance of your AuthController
    authController := controllers.NewAuthController()

    // Define your routes
    r.POST("/login", authController.Login)

    // Start the server
    if err := r.Run("192.168.23.53:8080"); err != nil {
        panic(err) // Handle error in starting the server
    }
}
