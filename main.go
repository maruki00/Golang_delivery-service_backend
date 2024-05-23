package main

import (
	routes "delivery/Services/Auth/Application/Routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()
	routes.AuthRouter(router)
	router.Run(":3000")

}
