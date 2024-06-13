package routes

import (
	"one_pay/config"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, appConfig *config.Config) {

	userGroup := router.Group("/user")
	transGroup := router.Group("/transaction")

	InitializeUserRoutes(userGroup, appConfig.DB)
	InitializeTransRoutes(transGroup, appConfig.DB)

}
