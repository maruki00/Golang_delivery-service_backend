package prouduct_routes

import (
	product_usergetway_controllers "delivery/Services/Product/UserGetway/Controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ProductRouter = func(router *gin.Engine, db *gorm.DB) {

	controller := product_usergetway_controllers.NewProductController(db)

	prefix := router.Group("/api/product")
	_ = prefix.POST("/insert", controller.Insert)
	_ = prefix.POST("/update", controller.Update)
	_ = prefix.DELETE("/delete", controller.Delete)
	_ = prefix.POST("/search", controller.Search)
	// _ = prefix.POST("/get", controller.Insert)
}
