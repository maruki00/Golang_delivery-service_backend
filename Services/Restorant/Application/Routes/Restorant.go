package restorant_application_routes

import "github.com/gin-gonic/gin"

func RestorantRouter(router *gin.Engine) {
	restorant := router.Group("/restaurant")

	_ = restorant.POST("/open", nil)
	_ = restorant.POST("/close", nil)

	menu := restorant.Group("/menu")
	menu.POST("/add", nil)
	menu.POST("/delete", nil)
	menu.POST("/edit", nil)
	menu.POST("/get", nil)
	menu.POST("/search", nil)
 
}
