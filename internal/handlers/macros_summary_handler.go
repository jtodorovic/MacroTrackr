package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/internal/models"
	"github.com/jtodorovic/macrotrackr/internal/services"
)

// GET /summary/:userId/today
func GetTodaysSummary(context *gin.Context) {
	var todaysSummary models.MacrosSummary

	userID := context.GetInt64("userID")

	dailyGoal, err := services.GetLatestGoalForUser(userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching user's daily goal."})
		return
	}

	todaysFoodLogs, err := services.FindFoodLogsByUserAndDate(userID, time.Now())
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching user's food logs."})
		return
	}

	var caloriesConsumed int32 = 0
	var carbsConsumed int32 = 0
	var proteinConsumed int32 = 0
	var fatsConsumed int32 = 0

	for _, fl := range todaysFoodLogs {
		fmt.Print(fl)
		caloriesConsumed += fl.Calories
		carbsConsumed += fl.Carbs
		proteinConsumed += fl.Protein
		fatsConsumed += fl.Fats
	}

	todaysSummary = models.MacrosSummary{
		UserID:           userID,
		Date:             time.Now(),
		CaloriesTarget:   dailyGoal.CaloriesTarget,
		ProteinTarget:    dailyGoal.ProteinTarget,
		CarbsTarget:      dailyGoal.CarbsTarget,
		FatsTarget:       dailyGoal.FatsTarget,
		CaloriesConsumed: caloriesConsumed,
		ProteinConsumed:  proteinConsumed,
		CarbsConsumed:    carbsConsumed,
		FatsConsumed:     fatsConsumed,
	}

	summaryResponse := services.ToSummaryResponse(todaysSummary)

	context.JSON(http.StatusOK, gin.H{"message": "Today's summary returned.", "todaysSummary": summaryResponse})
}
