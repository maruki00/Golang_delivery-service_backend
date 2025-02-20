package routes

import (
	"delivery/internal/auth/userGateWay/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var AuthRouter = func(router *gin.Engine, db *gorm.DB) {

	// repo := auth_infra_repository.NewAuthRepository(shared_configs.GetConfig())
	controller := controllers.NewAuthController(db)

	prefix := router.Group("/api/auth")
	_ = prefix.POST("/login", controller.Login)
	_ = prefix.POST("/register", controller.Register)
	_ = prefix.PATCH("/2f-confirm", controller.TwoFactoryConfirm)
	_ = prefix.PATCH("/logout", controller.Logout)

}
