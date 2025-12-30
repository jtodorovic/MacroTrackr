package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/models"
)

// GET /api/users
func GetUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		fmt.Println(err)

		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching users."})
		return
	}

	context.JSON(http.StatusOK, users)
}

// POST /api/users
func CreateUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request body."})
		return
	}

	user.ID = 1

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created.", "user": user})
}
