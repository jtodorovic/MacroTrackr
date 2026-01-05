package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/internal/models"
	"github.com/jtodorovic/macrotrackr/internal/services"
	"github.com/jtodorovic/macrotrackr/pkg/auth"
)

// POST /signup
func SignUp(context *gin.Context) {
	var req models.CreateUserRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request body."})
		return
	}

	user, err := services.CreateUser(req)
	if err != nil {
		fmt.Print(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created.",
		"user":    services.NewUserResponse(*user),
	})
}

// POST /login
func Login(context *gin.Context) {
	var credentials models.UserCredentials

	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request body."})
		return
	}

	if err := credentials.Validate(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	token, err := auth.GenerateToken(credentials.Email, credentials.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
