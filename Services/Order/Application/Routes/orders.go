package routes

import (
	"github.com/gin-gonic/gin"
)

var OrderRouter = func(router *gin.Engine) {
	order := router.Group("/order/")
	_ = order.POST("/create", nil)
	_ = order.POST("/cancel", nil)
	_ = order.POST("/history/{id}", nil)
	_ = order.POST("/pickup-confirm/{id}", nil)
	_ = router.POST("/chat", nil)
}
