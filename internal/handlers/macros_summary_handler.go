package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
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

func GetTodaysSummaryLambda(context context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	userIDStr := event.PathParameters["userId"]
	if userIDStr == "" {
		return events.APIGatewayV2HTTPResponse{StatusCode: 400, Body: `{"message":"userId required"}`}, nil
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: 400, Body: `{"message":"invalid userId"}`}, nil
	}

	// Fetch daily goal
	dailyGoal, err := services.GetLatestGoalForUserDynamoDB(userID)
	if err != nil {
		log.Printf("GetDailyGoal error: %+v", err)
		return events.APIGatewayV2HTTPResponse{StatusCode: 500, Body: `{"message":"Error fetching user's daily goal."}`}, nil
	}

	// Fetch today’s logs
	today := time.Now().Format("2006-01-02")
	todaysFoodLogs, err := services.FindFoodLogsByUserAndDateDynamoDB(userID, today)
	if err != nil {
		log.Printf("GetFoodLogs error: %+v", err)
		return events.APIGatewayV2HTTPResponse{StatusCode: 500, Body: `{"message":"Error fetching user's food logs."}`}, nil
	}

	var caloriesConsumed, carbsConsumed, proteinConsumed, fatsConsumed int32
	for _, fl := range todaysFoodLogs {
		caloriesConsumed += fl.Calories
		carbsConsumed += fl.Carbs
		proteinConsumed += fl.Protein
		fatsConsumed += fl.Fats
	}

	summary := models.MacrosSummary{
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

	respBody, _ := json.Marshal(map[string]interface{}{
		"message":       "Today's summary returned.",
		"todaysSummary": summary,
	})

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(respBody),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}
