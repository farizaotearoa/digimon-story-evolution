package controllers

import (
	"digimon-story-evolution/dto/request"
	"digimon-story-evolution/services"
	"digimon-story-evolution/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetDigimonDetails(c *gin.Context) {
	var req request.DigimonDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Fatal("Failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	details, err := services.GetDigimonDetails(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}

func GetDigimonEvolutions(c *gin.Context) {
	var req request.DigimonDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Fatal("Failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	details, err := services.GetDigimonEvolution(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}
