package shop_application_routes

import "github.com/gin-gonic/gin"

func RestorantRouter(router *gin.Engine) {
	shop := router.Group("/restaurant")

	_ = shop.POST("/open", nil)
	_ = shop.POST("/close", nil)

	menu := shop.Group("/menu")
	menu.POST("/add", nil)
	menu.POST("/delete", nil)
	menu.POST("/edit", nil)
	menu.POST("/get", nil)
	menu.POST("/search", nil)

}
