package main

import (
	"digimon-story-wiki/routes"
	"digimon-story-wiki/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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
	routes.SetupRoutes(r)

	// Start server
	appName := utils.Config.GetString(utils.AppsName)
	utils.Logger.Info(fmt.Sprintf("Starting %s", appName))

	r.Run(":9706")
}
