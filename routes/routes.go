package routes

import (
	"tryg/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.PUT("/value", controllers.CreateEntry)
		api.GET("/value/at", controllers.GetEntryByTimestamp)
	}
}