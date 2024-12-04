package controllers

import (
	"net/http"
	"travel-api/config"
	"travel-api/models"
	"github.com/gin-gonic/gin"
)

func CreateDestination(c *gin.Context) {
	var destination models.Destination
	if err := c.ShouldBindJSON(&destination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&destination)
	c.JSON(http.StatusOK, destination)
}

func GetDestinations(c *gin.Context) {
	var destinations []models.Destination
	config.DB.Preload("Reviews").Find(&destinations)
	c.JSON(http.StatusOK, destinations)
}

func UpdateDestination(c *gin.Context) {
	id := c.Param("id")
	var destination models.Destination
	if err := config.DB.First(&destination, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}
	if err := c.ShouldBindJSON(&destination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&destination)
	c.JSON(http.StatusOK, destination)
}

func DeleteDestination(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Destination{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
