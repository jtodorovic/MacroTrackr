package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/handlers"
	"github.com/jtodorovic/macrotrackr/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	// Auth & User Management
	server.POST("/signup", handlers.SignUp) // DONE
	server.POST("/login", handlers.Login)   // DONE
	// server.PUT("/users/:id", handlers.EditUser) // TODO

	// grouping authenticated routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// Daily Macros Goals
	authenticated.POST("/daily-goals", handlers.CreateDailyGoal)
	authenticated.GET("/daily-goals/latest", handlers.GetLatestGoalForUser)
	authenticated.PUT("/daily-goals/:id", handlers.EditDailyGoal)
	authenticated.DELETE("/daily-goals/:id", handlers.DeleteDailyGoal)

	// Entering Daily Meals - Food Logs
	authenticated.POST("/food-logs", handlers.CreateFoodLog) // TODO odavde menjati DTOove
	authenticated.PUT("/food-logs/:id", handlers.EditFoodLog)
	authenticated.DELETE("/food-log/:id", handlers.DeleteFoodLog)

	// Macros Intake Summary
	authenticated.GET("/summary/today", handlers.GetTodaysSummary)
	// api.GET("/summary/week", handlers.GetWeeklySummary) // cron job
	// api.GET("/summary/month", handlers.GetMonthlySummary) // cron job
}
