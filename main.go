package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Testimony struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Photo   string `json:"photo"`
	Content string `json:"content"`
}

type Destination struct {
	ID      uint   `json:"id"`
	Photo   string `json:"photo"`
	Location string `json:"location"`
	Price   float64 `json:"price"`
}

var DB *gorm.DB

func main() {
	var err error
	databasePath := os.Getenv("DATABASE_PATH")
	if databasePath == "" {
		databasePath = "./data/travel.db" // Caminho padrão
	}

	DB, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados!")
	}

	DB.AutoMigrate(&Testimony{}, &Destination{})

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Bem-vindo à API de Viagens!"})
	})

	router.GET("/depoimentos", getTestimonies)
	router.POST("/depoimentos", createTestimony)
	router.PUT("/depoimentos/:id", updateTestimony)
	router.DELETE("/depoimentos/:id", deleteTestimony)

	router.GET("/destinos", getDestinations)
	router.POST("/destinos", createDestination)
	router.PUT("/destinos/:id", updateDestination)
	router.DELETE("/destinos/:id", deleteDestination)

	router.GET("/depoimentos-home", getRandomTestimonies)

	for _, route := range router.Routes() {
		fmt.Println("Rota:", route.Path, "Método:", route.Method)
	}

	router.Run(":8080")
}

func getTestimonies(c *gin.Context) {
	var testimonies []Testimony
	DB.Find(&testimonies)
	c.JSON(200, testimonies)
}

func createTestimony(c *gin.Context) {
	var testimony Testimony
	if err := c.ShouldBindJSON(&testimony); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&testimony)
	c.JSON(201, testimony)
}

func updateTestimony(c *gin.Context) {
	var testimony Testimony
	id := c.Param("id")
	if err := DB.First(&testimony, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Depoimento não encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&testimony); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&testimony)
	c.JSON(200, testimony)
}

func deleteTestimony(c *gin.Context) {
	var testimony Testimony
	id := c.Param("id")
	if err := DB.First(&testimony, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Depoimento não encontrado"})
		return
	}
	DB.Delete(&testimony)
	c.Status(204)
}

func getDestinations(c *gin.Context) {
	var destinations []Destination
	DB.Find(&destinations)
	c.JSON(200, destinations)
}

func createDestination(c *gin.Context) {
	var destination Destination
	if err := c.ShouldBindJSON(&destination); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&destination)
	c.JSON(201, destination)
}

func updateDestination(c *gin.Context) {
	var destination Destination
	id := c.Param("id")
	if err := DB.First(&destination, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Destino não encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&destination); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&destination)
	c.JSON(200, destination)
}

func deleteDestination(c *gin.Context) {
	var destination Destination
	id := c.Param("id")
	if err := DB.First(&destination, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Destino não encontrado"})
		return
	}
	DB.Delete(&destination)
	c.Status(204)
}

func getRandomTestimonies(c *gin.Context) {
	var testimonies []Testimony
	DB.Order("RANDOM()").Limit(3).Find(&testimonies)
	c.JSON(200, testimonies)
}
