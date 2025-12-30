package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/config"
	"github.com/jtodorovic/macrotrackr/db"
	"github.com/jtodorovic/macrotrackr/routes"
)

func main() {

	config.LoadEnv()

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	log.Println("Server running on http://localhost:8080")
	server.Run()
}
