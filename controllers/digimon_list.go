package controllers

import (
	"digimon-story-evolution/dto/request"
	"digimon-story-evolution/services"
	"digimon-story-evolution/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDigimonList(c *gin.Context) {
	var req request.DigimonListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("Failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if req.SortBy == "" {
		req.SortBy = "number"
	}
	if req.SortOrder == "" {
		req.SortOrder = "asc"
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}

	digimon, err := services.GetAllDigimonList(c.Request.Context(), req)
	if err != nil {
		utils.Logger.Error("Failed to get digimon list", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, digimon)
}

func GetDigimonListSize(c *gin.Context) {
	var req request.DigimonListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("Failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	size, err := services.GetAllDigimonListSize(c.Request.Context(), req)
	if err != nil {
		utils.Logger.Error("Failed to get digimon list size", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"size": size})
}
