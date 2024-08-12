package auth_routes

import (
	Authcontrollers "delivery/Services/Auth/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
)

var AuthRouter = func(router *gin.Engine) {
	prefix := router.Group("/api/auth")
	_ = prefix.POST("/login", Authcontrollers.Login)
	_ = prefix.POST("/forget-password", Authcontrollers.Register)
	_ = prefix.POST("/2f-confirm", Authcontrollers.Register)
	_ = prefix.POST("/check", Authcontrollers.Register)
}
