package main

import (
	"digimon-story-evolution/routes"
	"digimon-story-evolution/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//Initialize Viper ConfigInterface
	if err := utils.InitConfig(); err != nil {
		panic("Failed to initialize configuration: " + err.Error())
	}

	//Initialize Logger
	utils.InitLogger()
	defer utils.Logger.Sync()

	//Connect To PostgreSQL
	utils.ConnectDatabase()

	// Set up routes
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins: utils.Config.GetStringSlice(utils.CorsAllowOrigins),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	routes.SetupRoutes(r)
	routes.SetupImagesRoutes(r)

	// Start server
	appName := utils.Config.GetString(utils.AppsName)
	utils.Logger.Info(fmt.Sprintf("Starting %s", appName))

	port := utils.Config.GetString(utils.AppsPort)
	if port == "" {
		port = os.Getenv("PORT")
	}
	r.Run(":" + port)
}
