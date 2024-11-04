// utils/jwt.go
package utils

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_secret_key") // Change this to a more secure key

// GenerateToken generates a JWT token
func GenerateToken(username string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateToken validates the JWT token
func ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
        }
        return jwtSecret, nil
    })
}
