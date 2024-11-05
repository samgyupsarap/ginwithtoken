package main

import (
	"gin-app/controllers" // Adjust the import path according to your project structure
	"gin-app/db"          // Import your db package for the database connection

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    // Connect to the database
    db.ConnectDB() // Ensure your db package has a Connect function to initialize the DB connection

    // Initialize Gin
    r := gin.Default()

        // CORS configuration
        r.Use(cors.New(cors.Config{
            AllowOrigins:     []string{"http://localhost:5173"}, // Allow your frontend origin
            AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
            AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
            ExposeHeaders:    []string{"Content-Length"},
            AllowCredentials: true,
        }))

    // Create an instance of your AuthController
    authController := controllers.NewAuthController()

    // Define your routes
    r.POST("/login", authController.Login)
    r.POST("/verify", controllers.VerifyIDToken)
    // Start the server
    if err := r.Run("192.168.23.53:8080"); err != nil {
        panic(err) // Handle error in starting the server
    }
}
