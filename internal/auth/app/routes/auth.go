package routes

import (
	auth_usergetway_controllers "delivery/internal/auth/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var AuthRouter = func(router *gin.Engine, db *gorm.DB) {

	// repo := auth_infrastructure_repository.NewAuthRepository(shared_configs.GetConfig())
	controller := auth_usergetway_controllers.NewAuthController(db)

	prefix := router.Group("/api/auth")
	_ = prefix.POST("/login", controller.Login)
	_ = prefix.POST("/register", controller.Register)
	_ = prefix.PATCH("/2f-confirm", controller.TwoFactoryConfirm)
	_ = prefix.PATCH("/logout", controller.Logout)

}
