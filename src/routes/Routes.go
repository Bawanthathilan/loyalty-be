package routes

import (
	"loyalty-be/src/controllers"

	"loyalty-be/utils"

	"github.com/gin-gonic/gin"
)



func SetupRoutes(router *gin.RouterGroup) {
	// Login
    router.POST("/login", controllers.Login)

    // Earn Point
    router.POST("/loyalty/accumulate", utils.RequireSession(), controllers.Earn)

    // Redeem Point
    router.POST("/loyalty/rewards/:reward_id/redeem", utils.RequireSession(), controllers.Redeem)

    // Balance
    router.GET("/loyalty/balance", utils.RequireSession(), controllers.Balance)

    // History
    router.GET("/loyalty/history",utils.RequireSession(), controllers.History)
}
