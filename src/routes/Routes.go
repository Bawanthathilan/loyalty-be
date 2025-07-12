package routes

import (
	"loyalty-be/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
	// Login
    router.GET("/login/:account_id", controllers.Login)

    // Search Accounts
    router.POST("/search", controllers.Search)

    // Earn Point
    router.POST("/loyalty/:account_id/accumulate", controllers.Earn)

    // Redeem Point
    router.POST("/loyalty/:account_id/redeem", controllers.Redeem)

    // Balance
    router.GET("/loyalty/:account_id/balance", controllers.Balance)

    // History
    router.GET("/loyalty/:account_id/history", controllers.History)
}
