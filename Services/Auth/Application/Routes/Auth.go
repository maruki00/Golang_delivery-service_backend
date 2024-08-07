package routes

import (
	Authcontrollers "delivery/Services/Auth/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
)

var AuthRouter = func(router *gin.Engine) {
	_ = router.POST("/login", Authcontrollers.Login)
	_ = router.POST("/register", Authcontrollers.Register)
}
