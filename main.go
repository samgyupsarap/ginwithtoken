package main

import (
	"gin-app/db"     // Import your db package for the database connection
	"gin-app/routes" // Import the routes package

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db.ConnectDB() // Ensure your db package has a ConnectDB function to initialize the DB connection

	// Initialize Gin with CORS configuration
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Initialize routes
	routes.SetupRoutes(r)

	// Start the server
	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err) // Handle error in starting the server
	}
}
