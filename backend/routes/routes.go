package routes

import (
	"digimon-story-wiki/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/digimon", controllers.GetDigimonList)
}
