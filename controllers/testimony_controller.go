package controllers

import (
	"net/http"
	"travel-api/config"
	"travel-api/models"

	"github.com/gin-gonic/gin"
)

func CreateTestimony(c *gin.Context) {
	var testimony models.Testimony
	if err := c.ShouldBindJSON(&testimony); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&testimony)
	c.JSON(http.StatusOK, testimony)
}

func GetTestimonies(c *gin.Context) {
	var testimonies []models.Testimony
	config.DB.Find(&testimonies)
	c.JSON(http.StatusOK, testimonies)
}

func UpdateTestimony(c *gin.Context) {
	id := c.Param("id")
	var testimony models.Testimony
	if err := config.DB.First(&testimony, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimony not found"})
		return
	}
	if err := c.ShouldBindJSON(&testimony); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&testimony)
	c.JSON(http.StatusOK, testimony)
}

func DeleteTestimony(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Testimony{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimony not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

func GetRandomTestimonies(c *gin.Context) {
	var testimonies []models.Testimony
	config.DB.Order("RANDOM()").Limit(3).Find(&testimonies)
	c.JSON(http.StatusOK, testimonies)
}
