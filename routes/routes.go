package routes

import (
	"gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/users", controllers.GetUsers)
}
