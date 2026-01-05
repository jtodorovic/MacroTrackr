package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/internal/models"
	"github.com/jtodorovic/macrotrackr/internal/services"
)

// GET /daily-goals/:user_id
func GetLatestGoalForUser(context *gin.Context) {
	userID := context.GetInt64("userID")

	goal, err := services.GetLatestGoalForUser(userID)

	// maybe return 404 TODO
	if err != nil {
		fmt.Print(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error fetching goal for user with ID %d", userID)})
		return
	}

	context.JSON(http.StatusOK, services.NewDailyGoalResponse(*goal))
}

// POST /daily-goals
func CreateDailyGoal(context *gin.Context) {
	var req models.CreateDailyGoalRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request body."})
		return
	}

	userID := context.GetInt64("userID")

	goal, err := services.CreateDailyGoal(userID, req)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating daily goal."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Daily goal created.", "dailyGoal": services.NewDailyGoalResponse(*goal)})
}

// PUT /daily-goals/:id
func EditDailyGoal(context *gin.Context) {
	ID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Print(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing goal ID."})
		return
	}

	dailyGoal, err := services.GetDailyGoalByID(ID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Daily goal not found."})
		return
	}

	userID := context.GetInt64("userID")
	if dailyGoal.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update daily goal."})
	}

	var req models.UpdateDailyGoalRequest

	if err = context.ShouldBindJSON(&req); err != nil {
		fmt.Print(err.Error())
		context.JSON(http.StatusNotFound, gin.H{"message": "Error parsing request body."})
		return
	}

	dailyGoal.ApplyUpdate(req)

	if err := dailyGoal.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating daily goal."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Daily goal updated.", "dailyGoal": services.NewDailyGoalResponse(*dailyGoal)})
}

// DELETE /daily-goals/:id
func DeleteDailyGoal(context *gin.Context) {
	ID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing goal ID."})
		return
	}

	dailyGoal, err := services.GetDailyGoalByID(ID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Daily goal not found."})
		return
	}

	userID := context.GetInt64("userID")
	if dailyGoal.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete daily goal."})
	}

	if err = dailyGoal.Delete(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Daily goal could not be deleted."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Daily goal deleted."})
}
