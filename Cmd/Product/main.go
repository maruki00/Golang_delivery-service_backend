package main

import (
	auth_routes "delivery/Internal/Auth/Application/Routes"
	order_routes "delivery/Internal/Order/Application/Routes"
	prouduct_routes "delivery/Internal/Product/Application/Routes"
	shared_configs "delivery/Internal/Shared/Application/Configs"
	shareddb "delivery/Internal/Shared/Infrastructure/DB"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := shared_configs.GetConfig()
	db := shareddb.NewMysqlDB_GORM(config)
	router := gin.Default()
	go auth_routes.AuthRouter(router, db)
	prouduct_routes.ProductRouter(router, db)
	order_routes.OrderRouter(router, db)

	router.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
}
