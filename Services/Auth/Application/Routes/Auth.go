package auth_routes

import (
	Authcontrollers "delivery/Services/Auth/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
)

controller := &auth_usergetway_controllers.AuthControllers{

}

var AuthRouter = func(router *gin.Engine) {
	prefix := router.Group("/api/auth")
	_ = prefix.POST("/login", controller.Login)
	_ = prefix.POST("/forget-password", controller.Register)
	_ = prefix.POST("/2f-confirm", controller.Register)
	_ = prefix.POST("/check", controller.Register)
}
