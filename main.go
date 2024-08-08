package main

import (
	authRoutes "delivery/Services/Auth/Application/Routes"
	orderRoutes "delivery/Services/Order/Application/Routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()
	authRoutes.AuthRouter(router)
	orderRoutes.OrderRouter(router)

	router.Run(":3000")
}
