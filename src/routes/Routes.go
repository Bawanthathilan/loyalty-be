package routes

import (
	"loyalty-be/src/controllers"

	"loyalty-be/utils"

	"github.com/gin-gonic/gin"
)



func SetupRoutes(router *gin.RouterGroup) {
	// Login
    router.GET("/login/:account_id", controllers.Login)

    // Search Accounts
    router.POST("/search", controllers.Search)

    // Earn Point
    router.POST("/loyalty/:account_id/accumulate", utils.RequireSession(), controllers.Earn)

    // Redeem Point
    router.POST("/loyalty/rewards/:reward_id/redeem", utils.RequireSession(), controllers.Redeem)

    // Balance
    router.GET("/loyalty/:account_id/balance", utils.RequireSession(), controllers.Balance)

    // History
    router.GET("/loyalty/:account_id/history",utils.RequireSession(), controllers.History)
}
