package routes

import (
	"loyalty-be/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
	// Define your routes here--
	router.GET("/example", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	router.GET("/locations", controllers.GetLocations)
	router.POST("/searchloyalty", controllers.SearchLoyalty)
}
