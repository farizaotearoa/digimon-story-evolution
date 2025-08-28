package main

import (
	"digimon-story-evolution/routes"
	"digimon-story-evolution/utils"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration from environment variables
	utils.LoadConfig()

	// Initialize Logger
	utils.InitLogger()
	defer utils.Logger.Sync()

	// Connect To PostgreSQL
	utils.ConnectDatabase()

	// Set up routes
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins: utils.GlobalConfig.CORS.AllowOrigins,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	routes.SetupRoutes(r)
	routes.SetupImagesRoutes(r)

	// Start server
	appName := utils.GlobalConfig.App.Name
	utils.Logger.Info(fmt.Sprintf("Starting %s", appName))

	port := utils.GlobalConfig.App.Port
	if port == "" {
		port = os.Getenv("PORT")
	}
	r.Run(":" + port)
}
