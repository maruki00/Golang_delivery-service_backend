package main

import (
	shared_configs "delivery/Services/Shared/Application/Configs"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := shared_configs.GetConfig()
	fmt.Println(filepath.Dir("./"))
	router := gin.Default()
	// auth_routes.AuthRouter(router)

	router.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
}
