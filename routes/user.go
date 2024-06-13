package routes

import (
	"one_pay/config"
	"one_pay/controller"
	"one_pay/repository"
	"one_pay/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeUserRoutes(router *gin.RouterGroup, appConfig *gorm.DB) {
	var (
		userRepository repository.UserRepository = repository.NewUserRepo(config.AppConfig.DB)
		userService    services.UserService      = services.NewUserService(userRepository)
		userHandler    controller.UserController = controller.NewUserController(userService)
	)
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
}
