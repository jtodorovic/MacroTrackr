package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/models"
)

// GET /summary/:userId/today
func GetTodaysSummary(context *gin.Context) {
	var todaysSummary models.MacrosSummary

	userID := context.GetInt64("userID")

	// fetch daily goal for user
	dailyGoal, err := models.GetLatestGoalForUser(userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching user's daily goal."})
		return
	}

	// fetch food logs for today for this user
	todaysFoodLogs, err := models.FindFoodLogsByUserAndDate(userID, time.Now())
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching user's food logs."})
		return
	}

	var caloriesConsumed int32 = 0
	var carbsConsumed int32 = 0
	var proteinConsumed int32 = 0
	var fatsConsumed int32 = 0

	for _, fl := range todaysFoodLogs {
		caloriesConsumed += fl.Calories

	}

	// combine it into dailySummary
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

	summaryResponse := models.ToSummaryResponse(todaysSummary)

	context.JSON(http.StatusOK, gin.H{"message": "Today's summary returned.", "todaysSummary": summaryResponse})
}
