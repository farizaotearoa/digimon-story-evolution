package routes

import (
	"digimon-story-evolution/controllers"
	"digimon-story-evolution/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/digimon/list", controllers.GetDigimonList)
	router.POST("/digimon/list/size", controllers.GetDigimonListSize)
	router.POST("/digimon/details", controllers.GetDigimonDetails)
	router.POST("/digimon/evolutions", controllers.GetDigimonEvolutions)
}

func SetupImagesRoutes(router *gin.Engine) {
	if strings.Contains(utils.GlobalConfig.Images.Path, "http") {
		router.GET("/images/*imagePath", func(c *gin.Context) {
			imagePath := c.Param("imagePath")
			cdnURL := utils.GlobalConfig.Images.Path + imagePath
			c.Redirect(http.StatusMovedPermanently, cdnURL)
		})
	} else {
		router.Static("/images", utils.GlobalConfig.Images.Path)
	}
}
