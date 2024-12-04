package config

import (
	"travel-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	connStr := "host=localhost port=8080 user=root password=root dbname=root sslmode=disable"
	var err error
	DB, err = gorm.Open(sqlite.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}
	DB.AutoMigrate(&models.Destination{}, &models.Testimony{})
}
