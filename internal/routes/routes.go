package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/internal/handlers"
	"github.com/jtodorovic/macrotrackr/internal/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	// Auth & User Management
	server.POST("/signup", handlers.SignUp)
	server.POST("/login", handlers.Login)

	// grouping authenticated routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// Daily Macros Goals
	authenticated.POST("/daily-goals", handlers.CreateDailyGoal)
	authenticated.GET("/daily-goals/latest", handlers.GetLatestGoalForUser)
	authenticated.PUT("/daily-goals/:id", handlers.EditDailyGoal)
	authenticated.DELETE("/daily-goals/:id", handlers.DeleteDailyGoal)

	// Entering Daily Meals - Food Logs
	authenticated.POST("/food-logs", handlers.CreateFoodLog)
	authenticated.PUT("/food-logs/:id", handlers.EditFoodLog)
	authenticated.DELETE("/food-logs/:id", handlers.DeleteFoodLog)

	// Generating Recipe - AI
	authenticated.POST("/generate-recipe", handlers.RecipeHandler)

	// Macros Intake Summary
	// authenticated.GET("/summary/:userID/today", handlers.GetTodaysSummary)
	// api.GET("/summary/week", handlers.GetWeeklySummary) // TODO: cron job
	// api.GET("/summary/month", handlers.GetMonthlySummary) // TODO: cron job
}
