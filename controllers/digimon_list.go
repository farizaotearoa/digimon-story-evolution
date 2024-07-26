package controllers

import (
	"digimon-story-evolution/dto/request"
	"digimon-story-evolution/services"
	"digimon-story-evolution/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetDigimonList(c *gin.Context) {
	var req request.DigimonListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Fatal("Failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	digimon, err := services.GetAllDigimonList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, digimon)
}

func GetDigimonListSize(c *gin.Context) {
	var req request.DigimonListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Fatal("Failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	size, err := services.GetAllDigimonListSize(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"size": size})
}
