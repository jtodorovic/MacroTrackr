package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/models"
)

// GET /goals/:userId
func GetGoalsForUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing user ID."})
		return
	}

	goals, err := models.GetGoalsForUser(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error fetching goals for user with ID %d", userId)})
		return
	}

	context.JSON(http.StatusOK, goals)
}
