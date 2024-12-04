package routes

import (
	"travel-api/controllers"

	"github.com/gin-gonic/gin"
)

func DestinationRoutes(router *gin.Engine) {
	destinations := router.Group("/destinos")
	{
		destinations.POST("", controllers.CreateDestination)
		destinations.GET("", controllers.GetDestinations)
		destinations.PUT("/:id", controllers.UpdateDestination)
		destinations.DELETE("/:id", controllers.DeleteDestination)
	}
}
