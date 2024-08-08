package main

import (
	authroutes "delivery/Services/Auth/Application/Routes"
	orderroutes "delivery/Services/Order/Application/Routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()
	authroutes.AuthRouter(router)
	orderroutes.OrderRouter(router)
	router.Run(":3000")

}
