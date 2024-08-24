package main

import (
	auth_routes "delivery/Services/Auth/Application/Routes"
	prouduct_routes "delivery/Services/Product/Application/Routes"
	shared_configs "delivery/Services/Shared/Application/Configs"
	shareddb "delivery/Services/Shared/Infrastructure/DB"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := shared_configs.GetConfig()
	db := shareddb.NewMysqlDB_GORM(config)
	fmt.Println(filepath.Dir("./"))
	router := gin.Default()
	auth_routes.AuthRouter(router, db)
	prouduct_routes.ProductRouter(router, db)

	router.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
}
