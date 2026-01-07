package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/internal/config"
	"github.com/jtodorovic/macrotrackr/internal/db"
	"github.com/jtodorovic/macrotrackr/internal/routes"
)

func main() {

	config.LoadEnv()

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	log.Println("Server running on http://localhost:8080")
	server.Run()
}
