package routes

import (
	"digimon-story-evolution/controllers"
	"digimon-story-evolution/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/digimon/list", controllers.GetDigimonList)
	router.POST("/digimon/list/size", controllers.GetDigimonListSize)
	router.POST("/digimon/details", controllers.GetDigimonDetails)
	router.POST("/digimon/evolutions", controllers.GetDigimonEvolutions)
}

func SetupImagesRoutes(router *gin.Engine) {
	fmt.Println("images path:", utils.Config.GetString(utils.ImagesPath))
	fmt.Println("contain http:", strings.Contains(utils.Config.GetString(utils.ImagesPath), "http"))

	if strings.Contains(utils.Config.GetString(utils.ImagesPath), "http") {
		router.GET("/images/*imagePath", func(c *gin.Context) {
			imagePath := c.Param("imagePath")
			cdnURL := utils.Config.GetString(utils.ImagesPath) + imagePath
			c.Redirect(http.StatusMovedPermanently, cdnURL)
		})
	} else {
		router.Static("/images", utils.Config.GetString(utils.ImagesPath))
	}
}
