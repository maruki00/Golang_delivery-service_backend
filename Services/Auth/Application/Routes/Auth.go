package routes

import (
	Authcontrollers "delivery/Services/Auth/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
)

var AuthRouter = func(router *gin.Engine) {
	_ = router.POST("/", Authcontrollers.Login)
	_ = router.POST("/", Authcontrollers.Login)
}
