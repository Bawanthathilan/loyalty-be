package main

import (
	"loyalty-be/config"
	"loyalty-be/src/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init(){
	config.LoadEnv()
}

func main() {
	router := gin.Default()

	store:= cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("loyalty_session", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins, change this in production
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Square-Version"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	api := router.Group("/api/v1")
	routes.SetupRoutes(api)

	router.Run()
}
