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
	utils.Logger.Info(fmt.Sprintf("Starting %s", utils.Config.GetString(utils.AppsName)))
	r.Run(":9706")
}
