package auth_routes

import (
	auth_usergetway_controllers "delivery/Services/Auth/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
)

var AuthRouter = func(router *gin.Engine) {

	controller := &auth_usergetway_controllers.AuthController{}

	prefix := router.Group("/api/auth")
	_ = prefix.POST("/login", controller.Login)
	// _ = prefix.POST("/forget-password", controller.Register)
	// _ = prefix.POST("/2f-confirm", controller.Register)
	// _ = prefix.POST("/check", controller.Register)
}
