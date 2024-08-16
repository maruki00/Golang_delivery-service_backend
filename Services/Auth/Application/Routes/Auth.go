package auth_routes

import (
	auth_infrastructure_repository "delivery/Services/Auth/Infrastructure/Repositories"
	auth_usergetway_controllers "delivery/Services/Auth/UserGateway/Controllers"
	shared_configs "delivery/Services/Shared/Application/Configs"

	"github.com/gin-gonic/gin"
)

var AuthRouter = func(router *gin.Engine) {

	repo := auth_infrastructure_repository.NewAuthRepository(shared_configs.GetConfig())
	controller := auth_usergetway_controllers.NewAuthController()

	prefix := router.Group("/api/auth")
	_ = prefix.POST("/login", controller.Login)
	// _ = prefix.POST("/forget-password", controller.Register)
	// _ = prefix.POST("/2f-confirm", controller.Register)
	// _ = prefix.POST("/check", controller.Register)
}
