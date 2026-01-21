package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/internal/handlers"
	"github.com/jtodorovic/macrotrackr/internal/middlewares"
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
	authenticated.POST("/daily-goals", handlers.CreateDailyGoal)            // DONE
	authenticated.GET("/daily-goals/latest", handlers.GetLatestGoalForUser) // DONE
	authenticated.PUT("/daily-goals/:id", handlers.EditDailyGoal)           // DONE
	authenticated.DELETE("/daily-goals/:id", handlers.DeleteDailyGoal)      // DONE

	// Entering Daily Meals - Food Logs
	authenticated.POST("/food-logs", handlers.CreateFoodLog)       // DONE
	authenticated.PUT("/food-logs/:id", handlers.EditFoodLog)      // DONE
	authenticated.DELETE("/food-logs/:id", handlers.DeleteFoodLog) // DONE

	// Macros Intake Summary
	// authenticated.GET("/summary/:userID/today", handlers.GetTodaysSummary) // DONE
	// api.GET("/summary/week", handlers.GetWeeklySummary) // cron job
	// api.GET("/summary/month", handlers.GetMonthlySummary) // cron job
}
