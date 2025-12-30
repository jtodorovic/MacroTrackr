package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/models"
)

// POST /food-logs
func CreateFoodLog(context *gin.Context) {
	var foodLog models.FoodLog

	err := context.ShouldBindJSON(&foodLog)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request body."})
		return
	}

	userID := context.GetInt64("userID")
	foodLog.UserID = userID

	err = foodLog.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating food log."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Food log created.", "foodLog": foodLog})
}

// PUT /food-logs/:id
func EditFoodLog(context *gin.Context) {
	ID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing food ID."})
		return
	}

	foodLog, err := models.FindFoodLogByID(ID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Food log not found."})
		return
	}

	userID := context.GetInt64("userID")
	if foodLog.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update food log."})
	}

	var updatedFood models.FoodLog
	err = context.ShouldBindJSON(&updatedFood)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Error parsing request body."})
		return
	}

	updatedFood.ID = ID // has to be set because our req body doesn't contain ID

	err = updatedFood.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating food log."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Food log updated."})
}

// DELETE /food-logs/:id
func DeleteFoodLog(context *gin.Context) {
	ID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing food ID."})
		return
	}

	foodLog, err := models.FindFoodLogByID(ID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Food log not found."})
		return
	}

	userID := context.GetInt64("userID")
	if foodLog.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete food log."})
	}

	err = foodLog.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Food log could not be deleted."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Food log deleted."})
}
