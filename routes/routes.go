package routes

import (
    "gin-app/controllers" // Adjust the import path according to your project structure
    "github.com/gin-gonic/gin"
)

// SetupRoutes initializes the routes for the application
func SetupRoutes() *gin.Engine {
    r := gin.Default()

    // Create an instance of the AuthController
    authController := controllers.NewAuthController()

    // Define the login route
    r.POST("/login", authController.Login)

    return r
}
