package controllers

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// FirebaseAuthClient is a global variable for Firebase Auth client
var FirebaseAuthClient *auth.Client

// InitializeFirebase initializes Firebase with the service account
func InitializeFirebase() error {
    opt := option.WithCredentialsFile("../oauth-ea02b-firebase-adminsdk-an3zu-805230c444.json") // Update path to your service account file

    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        return err
    }

    FirebaseAuthClient, err = app.Auth(context.Background())
    if err != nil {
        return err
    }

    log.Println("Firebase initialized successfully!")
    return nil
}


// VerifyIDToken verifies a Google ID token sent from the client
func VerifyIDToken(c *gin.Context) {
    var req struct {
        IDToken string `json:"idToken"`
    }

    // Bind JSON from the client
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Verify the ID token with Firebase
    token, err := FirebaseAuthClient.VerifyIDToken(context.Background(), req.IDToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    // Successfully verified, get the user ID
    userID := token.UID
    c.JSON(http.StatusOK, gin.H{"message": "Token verified", "userID": userID})
}
