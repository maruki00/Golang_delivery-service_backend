package routes

import (
	Authcontrollers "delivery/Services/Auth/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
)

var AuthRouter = func(router *gin.Engine) {
	controller := &Authcontrollers.Authcontroller{}
	_ = router.GET("/", controller.Login)
}
