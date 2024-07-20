package controllers

import (
	"digimon-story-wiki/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDigimonList(c *gin.Context) {
	digimon, err := services.GetAllDigimonList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, digimon)
}
