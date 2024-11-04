package utils

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware checks for a valid JWT token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
            c.Abort()
            return
        }

        token, err := ValidateToken(tokenString)
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
