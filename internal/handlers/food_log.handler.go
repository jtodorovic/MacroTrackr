package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/internal/models"
	"github.com/jtodorovic/macrotrackr/internal/services"
)

// POST /food-logs
// factory function that returns handlerFunc
func CreateFoodLog(c *gin.Context) {
	var req models.CreateFoodLogFromAPIRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	req.Food = strings.TrimSpace(req.Food)
	if req.Food == "" || req.Weight <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "food and weight are required",
		})
		return
	}

	food, err := services.FetchNutrition(req.Food)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "nutrition service unavailable",
		})
		return
	}

	macros := models.ExtractMacros(food, req.Weight)

	c.JSON(http.StatusOK, gin.H{
		"food":     food.Description,
		"weight":   req.Weight,
		"calories": macros.Calories,
		"protein":  macros.Protein,
		"carbs":    macros.Carbs,
		"fat":      macros.Fats,
	})
}

// PUT /food-logs/:id
func EditFoodLog(context *gin.Context) {
	ID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing food ID."})
		return
	}

	foodLog, err := services.FindFoodLogByID(ID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Food log not found."})
		return
	}

	userID := context.GetInt64("userID")
	if foodLog.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update food log."})
	}

	var req models.UpdateFoodLogRequest

	if err = context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Error parsing request body."})
		return
	}

	foodLog.ApplyUpdate(req)

	if err := foodLog.Update(); err != nil {
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

	foodLog, err := services.FindFoodLogByID(ID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Food log not found."})
		return
	}

	userID := context.GetInt64("userID")
	if foodLog.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete food log."})
		return
	}

	if err = foodLog.Delete(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Food log could not be deleted."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Food log deleted."})
}
