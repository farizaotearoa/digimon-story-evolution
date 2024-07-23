package routes

import (
	"digimon-story-wiki/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/digimon/list", controllers.GetDigimonList)
	router.POST("/digimon/list/size", controllers.GetDigimonListSize)
	router.POST("/digimon/details", controllers.GetDigimonDetails)
	router.POST("/digimon/evolutions", controllers.GetDigimonEvolutions)
}
