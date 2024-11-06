package routes

import (
	"gin-app/controllers" // Adjust the import path according to your project structure

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the routes for the application
func SetupRoutes(r *gin.Engine) {
	// Create instances of AuthController and RegisterController
	authController := controllers.NewAuthController()
	registerController := controllers.NewRegisterController()

	// Define the login and register routes
	r.POST("/login", authController.Login)
	r.POST("/register", registerController.Register)

	// Add other routes here, e.g., token verification
	r.POST("/verify", controllers.VerifyIDToken)

    // Define the route to get all animals
    r.GET("/animals", controllers.GetAnimals)
}
