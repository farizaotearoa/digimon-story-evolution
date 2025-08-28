package controllers

import (
	"digimon-story-evolution/dto/request"
	"digimon-story-evolution/services"
	"digimon-story-evolution/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDigimonDetails(c *gin.Context) {
	var req request.DigimonDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("Failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	details, err := services.GetDigimonDetails(req)
	if err != nil {
		utils.Logger.Error("Failed to get digimon details", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, details)
}

func GetDigimonEvolutions(c *gin.Context) {
	var req request.DigimonDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("Failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	details, err := services.GetDigimonEvolution(req)
	if err != nil {
		utils.Logger.Error("Failed to get digimon evolutions", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, details)
}
