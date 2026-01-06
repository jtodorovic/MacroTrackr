package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/internal/config"
	"github.com/jtodorovic/macrotrackr/internal/db"
	"github.com/jtodorovic/macrotrackr/internal/routes"
)

func main() {

	config.LoadEnv()

	db.InitDB()

	nutritionAPIKey := os.Getenv("USDA_API_KEY")
	if nutritionAPIKey == "" {
		log.Fatal("USDA_API_KEY is not set")
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	log.Println("Server running on http://localhost:8080")
	server.Run()
}
