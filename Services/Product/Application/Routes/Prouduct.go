package prouduct_routes

import (
	auth_usergetway_controllers "delivery/Services/Auth/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ProductRouter = func(router *gin.Engine, db *gorm.DB) {

	// repo := auth_infrastructure_repository.NewAuthRepository(shared_configs.GetConfig())
	controller := auth_usergetway_controllers.NewProductController(db)

	prefix := router.Group("/api/product")
	_ = prefix.POST("/login", controller.Login)

}
