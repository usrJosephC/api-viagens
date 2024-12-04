package routes

import (
	"travel-api/controllers"

	"github.com/gin-gonic/gin"
)

func TestimonyRoutes(router *gin.Engine) {
	testimonies := router.Group("/depoimentos")
	{
		testimonies.POST("", controllers.CreateTestimony)
		testimonies.GET("", controllers.GetTestimonies)
		testimonies.PUT("/:id", controllers.UpdateTestimony)
		testimonies.DELETE("/:id", controllers.DeleteTestimony)
	}
	router.GET("/depoimentos-home", controllers.GetRandomTestimonies)
}
