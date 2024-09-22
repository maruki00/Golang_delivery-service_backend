package main

import (
	auth_routes "delivery/Services/Auth/Application/Routes"
	order_routes "delivery/Services/Order/Application/Routes"
	prouduct_routes "delivery/Services/Product/Application/Routes"
	shared_configs "delivery/Services/Shared/Application/Configs"
	shareddb "delivery/Services/Shared/Infrastructure/DB"
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
