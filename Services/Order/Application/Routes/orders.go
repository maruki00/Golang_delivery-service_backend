package routes

import (
	Authcontrollers "delivery/Services/Auth/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
)

var AuthRouter = func(router *gin.Engine) {
	_ = router.POST("/login", Authcontrollers.Login)
	_ = router.POST("/register", Authcontrollers.Register)
	order := router.Group("/order/")
	_ = order.POST("/create", Authcontrollers.Register)
	_ = order.POST("/cancel", Authcontrollers.Register)
	_ = order.POST("/history/{id}", Authcontrollers.Register)
	_ = order.POST("/pickup-confirm/{id}", Authcontrollers.Register)
	_ = router.POST("/chat", Authcontrollers.Register)
}
