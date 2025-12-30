package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/handlers"
)

func RegisterRoutes(server *gin.Engine) {
	api := server.Group("/api")

	api.GET("/users", handlers.GetUsers)
	api.POST("/users", handlers.CreateUser)
	api.GET("/goals/:userId", handlers.GetGoalsForUser)
}
