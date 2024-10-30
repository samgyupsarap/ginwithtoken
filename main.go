package main

import (
	"gin-app/db"     // Import the db package
	"gin-app/routes" // Import your routes package
	"github.com/gin-gonic/gin"
)

func main() {
    db.ConnectDB() // Initialize the database connection

    r := gin.Default()

    // Setup routes
    routes.SetupRoutes(r)

    // Start the server
    r.Run(":8080") // Start the server on port 8080
}
