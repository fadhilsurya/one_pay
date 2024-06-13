package routes

import (
	"one_pay/config"
	"one_pay/controller"
	"one_pay/middleware"
	"one_pay/repository"
	"one_pay/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeTransRoutes(router *gin.RouterGroup, appConfig *gorm.DB) {
	var (
		transRepository repository.TransactionRepo  = repository.NewTransactionRepo(config.AppConfig.DB)
		userRepository  repository.UserRepository   = repository.NewUserRepo(config.AppConfig.DB)
		transService    services.TransactionService = services.NewTransactionService(transRepository, userRepository)
		transHandler    controller.TransController  = controller.NewTransController(transService)
	)
	router.POST("/", middleware.AuthMiddleware(), transHandler.Create)
	router.GET("/", middleware.AuthMiddleware(), transHandler.GetTransactionHistory)
}
